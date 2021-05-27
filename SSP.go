package ssp

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

var dialect = ""

// Data is a line in map that link the database field with datatable field
type Data struct {
	Db        string                                                                  //name of column
	Dt        interface{}                                                             //id of column in client (int or string)
	Cs        bool                                                                    //case sensitive - optional default false
	Formatter func(data interface{}, row map[string]interface{}) (interface{}, error) // - optional
}

// MessageDataTable is theresponse object
type MessageDataTable struct {
	Draw            int           `json:"draw"`
	RecordsTotal    int           `json:"recordsTotal"`
	RecordsFiltered int           `json:"recordsFiltered"`
	Data            []interface{} `json:"data,nilasempty"`
}

// Controller emulate the beego controller
type Controller interface {
	GetString(string, ...string) string
}

// Simple is a main method, externally called
func Simple(c Controller, conn *gorm.DB,
	table string,
	columns []Data) (responseJSON MessageDataTable, err error) {

	dialect = conn.Dialect().GetName()

	responseJSON.Draw = DrawNumber(c)
	dbConfig(conn)

	columnsType, err := initBinding(conn, "*", table, make(map[string]string, 0))

	// Build the SQL query string from the request
	rows, err := conn.Select("*").
		Scopes(limit(c),
			filterGlobal(c, columns, columnsType),
			filterIndividual(c, columns, columnsType),
			order(c, columns)).
		Table(table).
		Rows()
	if err != nil {
		return
	}

	responseJSON.Data, err = dataOutput(columns, rows)
	if err != nil {
		return
	}

	//search in DDBB recordsFiltered
	err = conn.Scopes(filterGlobal(c, columns, columnsType),
		filterIndividual(c, columns, columnsType)).
		Table(table).
		Count(&responseJSON.RecordsFiltered).Error
	if err != nil {
		return
	}

	//search in DDBB recordsTotal
	err = conn.Table(table).Count(&responseJSON.RecordsTotal).Error
	if err != nil {
		return
	}

	defer rows.Close()
	return
}

// Complex is a main method, externally called
func Complex(c Controller, conn *gorm.DB, table string, columns []Data,
	whereResult []string,
	whereAll []string,
	whereJoin map[string]string) (responseJSON MessageDataTable, err error) {

	dialect = conn.Dialect().GetName()

	responseJSON.Draw = DrawNumber(c)
	dbConfig(conn)

	// Build the SQL query string from the request
	whereResultFlated := flated(whereResult)
	whereAllFlated := flated(whereAll)

	selectQuery, err := buildSelect(table, whereJoin, conn)
	if err != nil {
		return
	}

	columnsType, err := initBinding(conn, selectQuery, table, whereJoin)
	if err != nil {
		return
	}

	rows, err := conn.Select(selectQuery).
		Scopes(
			setJoins(whereJoin),
			limit(c),
			filterGlobal(c, columns, columnsType),
			filterIndividual(c, columns, columnsType),
			order(c, columns)).
		Where(whereResultFlated).
		Where(whereAllFlated).
		Table(table).
		Rows()
	if err != nil {
		return
	}

	responseJSON.Data, err = dataOutput(columns, rows)
	if err != nil {
		return
	}

	//search in DDBB recordsFiltered
	err = conn.
		Scopes(
			setJoins(whereJoin),
			filterGlobal(c, columns, columnsType),
			filterIndividual(c, columns, columnsType)).
		Where(whereResultFlated).
		Where(whereAllFlated).
		Table(table).
		Count(&responseJSON.RecordsFiltered).Error
	if err != nil {
		return
	}

	//search in DDBB recordsTotal
	err = conn.Table(table).
		Scopes(setJoins(whereJoin)).
		Where(whereAllFlated).Count(&responseJSON.RecordsTotal).Error
	if err != nil {
		return
	}

	defer rows.Close()
	return
}

func dataOutput(columns []Data, rows *sql.Rows) ([]interface{}, error) {
	out := make([]interface{}, 0)

	for rows.Next() {
		fields, err := getFields(rows)
		if err != nil {
			return nil, err
		}

		row := make(map[string]interface{})

		for j := 0; j < len(columns); j++ {
			column := columns[j]
			var dt string
			if column.Dt == nil {
				return nil, fmt.Errorf("Dt cannot be nil in column[%v]", j)
			}

			vType := reflect.TypeOf(column.Dt)
			if vType.String() == "string" {
				dt = column.Dt.(string)
			} else {
				dt = strconv.Itoa(column.Dt.(int))
			}

			db := column.Db
			// Is there a formatter?
			if column.Formatter != nil {
				var err error
				row[dt], err = column.Formatter(fields[db], fields)
				if err != nil {
					return nil, err
				}
			} else {
				row[dt] = fields[db]
			}

		}
		out = append(out, row)
	}

	return out, nil
}

func flated(whereArray []string) string {
	query := ""
	for _, where := range whereArray {
		if query != "" && where != "" {
			query += " AND "
		}
		query += where
	}
	return query
}

func buildSelect(table string, join map[string]string, conn *gorm.DB) (query string, err error) {
	query = fmt.Sprintf("%s.*", table)
	if len(join) == 0 {
		return
	}

	subQuery, err := addFieldsSelect(table, conn)
	query += subQuery
	for tableName := range join {
		subQuery, err = addFieldsSelect(tableName, conn)
		query += subQuery
	}
	return
}

func addFieldsSelect(table string, conn *gorm.DB) (query string, err error) {
	columnsType, err := initBinding(conn, "*", table, make(map[string]string, 0))
	for _, columnInfo := range columnsType {
		query += fmt.Sprintf(", %s.%s AS \"%s.%s\"", table, columnInfo.Name(), table, columnInfo.Name())
	}
	return
}

func setJoins(joins map[string]string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, join := range joins {
			db = db.Joins(join)
		}
		return db
	}
}

func setQuery(db *gorm.DB, query, param, logic string) *gorm.DB {
	if logic == "where" {
		if param == "" {
			db = db.Where(query)
		} else {
			db = db.Where(query, param)
		}
	} else {
		if param == "" {
			db = db.Or(query)
		} else {
			db = db.Or(query, param)
		}
	}

	return db
}

func setGlobalQuery(db *gorm.DB, query, param string, first bool) *gorm.DB {
	logic := "or"
	if first {
		logic = "where"
	}
	return setQuery(db, query, param, logic)
}

//database func
func filterGlobal(c Controller, columns []Data, columnsType []*sql.ColumnType) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		str := c.GetString("search[value]")
		if str == "" {
			return db
		}
		requestRegex := ParamToBool(c, "search[regex]")

		//all columns filtering
		var i int
		first := true
		for i = 0; ; i++ {
			keyColumnsI := fmt.Sprintf("columns[%d][data]", i)

			keyColumnsData := c.GetString(keyColumnsI)
			if keyColumnsData == "" {
				break
			}
			columnIdx := search(columns, keyColumnsData)

			requestColumnQuery := fmt.Sprintf("columns[%d][searchable]", i)
			requestColumn := c.GetString(requestColumnQuery)

			if columnIdx > -1 && requestColumn == "true" {

				query, param := bindingTypes(str, columnsType, columns[columnIdx], requestRegex)
				if query == "" {
					continue
				}
				db = setGlobalQuery(db, query, param, first)
				first = false

			} else {
				if columnIdx < 0 && requestColumn == "true" {
					fmt.Printf("(002) Do you forgot searchable: false in column %v ? or wrong column name in client side\n (client field data: must be same than server side DT: field)\n", keyColumnsData)
				}
			}
		}
		return db
	}
}

func filterIndividual(c Controller, columns []Data, columnsType []*sql.ColumnType) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Individual column filtering
		var i int
		for i = 0; ; i++ {
			keyColumnsI := fmt.Sprintf("columns[%d][data]", i)

			keyColumnsData := c.GetString(keyColumnsI)
			if keyColumnsData == "" {
				break
			}

			columnIdx := search(columns, keyColumnsData)

			requestColumnQuery := fmt.Sprintf("columns[%d][searchable]", i)
			requestColumn := c.GetString(requestColumnQuery)

			requestColumnQuery = fmt.Sprintf("columns[%d][search][value]", i)
			str := c.GetString(requestColumnQuery)
			if columnIdx > -1 && requestColumn == "true" && str != "" {
				requestRegexQuery := fmt.Sprintf("columns[%d][search][regex]", i)
				requestRegex := ParamToBool(c, requestRegexQuery)
				query, param := bindingTypes(str, columnsType, columns[columnIdx], requestRegex)
				if query == "" {
					continue
				}
				db = setQuery(db, query, param, "where")

			} else {
				if columnIdx < 0 && requestColumn == "true" {
					fmt.Printf("(001) Do you forgot searchable: false in column %v ? or wrong column name in client side\n (client field data: must be same than server side DT: field)\n", keyColumnsData)
				}
			}
		}
		return db
	}
}

//Refactor this
func order(c Controller, columns []Data) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if c.GetString("order[0][column]") != "" {
			var i int
			for i = 0; ; i++ {
				columnIdxTittle := fmt.Sprintf("order[%d][column]", i)

				columnIdxOrder := c.GetString(columnIdxTittle)
				if columnIdxOrder == "" {
					break
				}

				columnIdxTittle = fmt.Sprintf("columns[%s][data]", columnIdxOrder)
				requestColumnData := c.GetString(columnIdxTittle)
				columnIdx := search(columns, requestColumnData)

				columnIdxTittle = fmt.Sprintf("columns[%s][orderable]", columnIdxOrder)

				if columnIdx > -1 && c.GetString(columnIdxTittle) == "true" {

					column := columns[columnIdx]
					columnIdxTittle = fmt.Sprintf("order[%d][dir]", i)
					requestColumnData = c.GetString(columnIdxTittle)

					order := "desc"
					if requestColumnData == "asc" {
						order = "asc"
					}

					order = checkOrderDialect(order)

					query := fmt.Sprintf("%s %s", column.Db, order)
					db = db.Order(query)
				} else {
					if columnIdx < 0 && c.GetString(columnIdxTittle) == "true" {
						fmt.Printf("(003) Do you forgot orderable: false in column %v ?\n", columnIdxOrder)
					}
				}
			}
		}
		return db
	}
}
func checkOrderDialect(order string) string {
	if dialect == "sqlite3" {
		if order == "asc" {
			return "desc"
		}
		return "asc"
	}

	return order
}

func limit(c Controller) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		start, err := strconv.Atoi(c.GetString("start"))
		if err != nil || start < 0 {
			start = 0
		}

		length, err := strconv.Atoi(c.GetString("length"))

		if err != nil || length < 0 {
			length = 10
		}

		return db.Offset(start).Limit(length)
	}
}

func search(column []Data, keyColumnsI string) int {
	var i int
	for i = 0; i < len(column); i++ {
		data := column[i]
		if data.Dt == nil {
			continue
		}
		var field string
		vType := reflect.TypeOf(data.Dt)
		if vType.String() == "string" {
			field = data.Dt.(string)
		} else {
			field = strconv.Itoa(data.Dt.(int))
		}

		if field == keyColumnsI {
			return i
		}
	}
	return -1
}

//check if searchable field is string
func bindingTypes(value string, columnsType []*sql.ColumnType, column Data, isRegEx bool) (string, string) {
	columndb := column.Db
	for _, columnInfo := range columnsType {
		if columnInfo.Name() == columndb {

			searching := columnInfo.DatabaseTypeName()
			if strings.Contains(searching, "varchar") {
				searching = "varchar"
			}
			return bindingTypesQuery(searching, CheckReserved(columndb), value, columnInfo, isRegEx, column)

		}
	}

	return "", ""
}

func bindingTypesQuery(searching, columndb, value string, columnInfo *sql.ColumnType, isRegEx bool, column Data) (string, string) {
	switch searching {
	case "string", "TEXT", "varchar", "VARCHAR":
		if isRegEx {
			return regExp(columndb, value)
		}

		if column.Cs {
			return fmt.Sprintf("%s LIKE ?", columndb), "%" + value + "%"
		}
		return fmt.Sprintf("Lower(%s) LIKE ?", columndb), "%" + strings.ToLower(value) + "%"
	case "UUID", "blob":
		if isRegEx {
			return regExp(fmt.Sprintf("CAST(%s AS TEXT)", columndb), value)
		}
		return fmt.Sprintf("%s = ?", columndb), value
	case "int32", "INT4", "INT8", "integer", "INTEGER", "bigint":
		if isRegEx {
			return regExp(fmt.Sprintf("CAST(%s AS TEXT)", columndb), value)
		}
		_, err := strconv.Atoi(value)
		if err != nil {
			return "", ""
		}
		return fmt.Sprintf("%s = ?", columndb), value
	case "bool", "BOOL":
		boolval, err := strconv.ParseBool(value)
		queryval := "NOT"
		if err == nil && boolval {
			queryval = ""
		}
		return fmt.Sprintf("%s IS %s TRUE", columndb, queryval), ""
	case "real", "NUMERIC":
		if isRegEx {
			return regExp(fmt.Sprintf("CAST(%s AS TEXT)", columndb), value)
		}
		fmt.Print("(005) GORMSSP WARNING: Serarching float values, float cannot be exactly equal\n")
		_, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return "", ""
		}
		return fmt.Sprintf("%s = ?", columndb), value
	default:
		fmt.Printf("(004) GORMSSP New type %v\n", columnInfo.DatabaseTypeName())
		return "", ""
	}
}

func regExp(columndb, value string) (string, string) {
	switch dialect {
	case "sqlite3":
		//TODO make regexp
		return fmt.Sprintf("Lower(%s) LIKE '%s'", columndb, "%"+strings.ToLower(value)+"%"), ""
	case "postgres":
		return fmt.Sprintf("%s ~* '%s'", columndb, value), ""
	default:
		return fmt.Sprintf("%s ~* '%s'", columndb, value), ""
	}
}

// https://github.com/jinzhu/gorm/issues/1167
func getFields(rows *sql.Rows) (map[string]interface{}, error) {

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	length := len(columns)
	current := makeResultReceiver(length)

	columnsType, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	err = rows.Scan(current...)
	if err != nil {
		return nil, err
	}

	value := make(map[string]interface{})
	for i := 0; i < length; i++ {
		key := columns[i]
		val := *(current[i]).(*interface{})
		if val == nil {
			value[key] = val
			continue
		}
		vType := reflect.TypeOf(val)
		searching := columnsType[i].DatabaseTypeName()
		if strings.Contains(searching, "varchar") {
			searching = "varchar"
		}
		value[key], err = getFieldsSearch(searching, key, val, vType)
		if err != nil {
			return nil, err
		}

	}
	return value, nil
}

func getFieldsSearch(searching, key string, val interface{}, vType reflect.Type) (interface{}, error) {
	switch searching {

	case "string", "TEXT", "varchar", "VARCHAR":
		return val.(string), nil
	case "int32", "INT4", "INT8", "integer", "bigint", "INTEGER":
		return val.(int64), nil
	case "NUMERIC", "real":
		switch vType.String() {
		case "[]uint8":
			return strconv.ParseFloat(string(val.([]uint8)), 64)
		case "float64":
			return val.(float64), nil
		default:
			return val, nil
		}
	case "bool", "BOOL":
		switch vType.String() {
		case "int64":
			return val.(int64) == 1, nil
		case "bool":
			return val.(bool), nil
		default:
			return val, nil
		}

	case "TIMESTAMPTZ", "datetime":
		return val.(time.Time), nil
	case "UUID", "blob":
		switch vType.String() {
		case "[]uint8":
			return string(val.([]uint8)), nil
		case "string":
			return val, nil
		default:
			return val, nil
		}
	case "BYTEA":
		return val.([]byte), nil
	default:
		fmt.Printf("(006) GORMSSP New type: %v for key: %v\n", searching, key)
		return val, nil
	}
}

func makeResultReceiver(length int) []interface{} {
	result := make([]interface{}, 0, length)
	for i := 0; i < length; i++ {
		var current interface{}
		current = struct{}{}
		result = append(result, &current)
	}
	return result
}

func initBinding(db *gorm.DB, selectQuery, table string, whereJoin map[string]string) ([]*sql.ColumnType, error) {
	rows, err := db.Select(selectQuery).
		Table(table).
		Scopes(
			setJoins(whereJoin),
		).
		Limit(0).
		Rows()
	if err != nil {
		return nil, err
	}

	columnsType, err := rows.ColumnTypes()

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return columnsType, nil
}

func dbConfig(conn *gorm.DB) {
	if dialect == "sqlite3" {
		conn.Exec("PRAGMA case_sensitive_like = ON;")
	}
}
