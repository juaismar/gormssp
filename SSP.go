package SSP

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type Data struct {
	Db        string
	Dt        interface{}
	Formatter func(data interface{}, row map[string]interface{}) interface{}
}

type MessageDataTable struct {
	Draw            int           `json:"draw"`
	RecordsTotal    int           `json:"recordsTotal"`
	RecordsFiltered int           `json:"recordsFiltered"`
	Data            []interface{} `json:"data,nilasempty"`
}

func Simple(c interface {
	GetString(string, ...string) string
}, conn *gorm.DB,
	table string,

	columns map[int]Data) MessageDataTable {

	draw := 0
	draw, err := strconv.Atoi(c.GetString("draw"))
	check(err)

	columnsType := initBinding(conn, table)

	// Build the SQL query string from the request
	rows, err := conn.Select("*").
		Scopes(limit(c),
			filterGlobal(c, columns, columnsType),
			filterIndividual(c, columns, columnsType),
			order(c, columns)).
		Table(table).
		Rows()

	check(err)

	Datas := dataOutput(columns, rows)

	//search in DDBB recordsFiltered
	var recordsFiltered int
	conn.Scopes(filterGlobal(c, columns, columnsType),
		filterIndividual(c, columns, columnsType)).
		Table(table).
		Count(&recordsFiltered)

	//search in DDBB recordsTotal
	var recordsTotal int
	conn.Table(table).Count(&recordsTotal)

	responseJSON := MessageDataTable{
		Draw:            draw,
		RecordsTotal:    recordsTotal,
		RecordsFiltered: recordsFiltered,
		Data:            Datas,
	}

	defer rows.Close()
	return responseJSON
}

func Complex(c interface {
	GetString(string, ...string) string
}, conn *gorm.DB, table string, columns map[int]Data,
	whereResult []string, whereAll []string) MessageDataTable {

	draw := 0
	draw, err := strconv.Atoi(c.GetString("draw"))
	check(err)

	columnsType := initBinding(conn, table)

	// Build the SQL query string from the request
	whereResultFlated := flated(whereResult)
	whereAllFlated := flated(whereAll)

	rows, err := conn.Select("*").
		Scopes(limit(c),
			filterGlobal(c, columns, columnsType),
			filterIndividual(c, columns, columnsType),
			order(c, columns)).
		Where(whereResultFlated).
		Where(whereAllFlated).
		Table(table).
		Rows()

	check(err)
	Datas := dataOutput(columns, rows)

	//search in DDBB recordsFiltered
	var recordsFiltered int
	conn.Scopes(filterGlobal(c, columns, columnsType),
		filterIndividual(c, columns, columnsType)).
		Where(whereResultFlated).
		Where(whereAllFlated).
		Table(table).
		Count(&recordsFiltered)

	//search in DDBB recordsTotal
	var recordsTotal int
	conn.Table(table).Count(&recordsTotal)

	responseJSON := MessageDataTable{
		Draw:            draw,
		RecordsTotal:    recordsTotal,
		RecordsFiltered: recordsFiltered,
		Data:            Datas,
	}

	defer rows.Close()
	return responseJSON
}

func dataOutput(columns map[int]Data, rows *sql.Rows) []interface{} {
	out := make([]interface{}, 0)

	for rows.Next() {
		fields := getFields(rows)

		row := make(map[string]interface{})

		for j := 0; j < len(columns); j++ {
			column := columns[j]
			var dt string

			vType := reflect.TypeOf(column.Dt)
			if vType.String() == "string" {
				dt = column.Dt.(string)
			} else {
				dt = strconv.Itoa(column.Dt.(int))
			}

			db := column.Db
			// Is there a formatter?
			if column.Formatter != nil {
				row[dt] = column.Formatter(fields[db], fields)
			} else {
				row[dt] = fields[db]
			}

		}
		out = append(out, row)
	}

	return out
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

//database func
func filterGlobal(c interface {
	GetString(string, ...string) string
}, columns map[int]Data, columnsType []*sql.ColumnType) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		globalSearch := ""
		str := c.GetString("search[value]")
		//all columns filtering
		if str != "" {
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

				if columnIdx > -1 && requestColumn == "true" {
					column := columns[columnIdx]
					columndb := column.Db

					query := bindingTypes(columndb, str, columnsType)

					if globalSearch != "" && query != "" {
						globalSearch += " OR "
					}

					globalSearch += query
				} else {
					if columnIdx < 0 && requestColumn == "true" {
						fmt.Printf("Do you forgot searchable: false in column %v ?\n", keyColumnsData)
					}
				}
			}
		}
		return db.Where(globalSearch)
	}
}

func filterIndividual(c interface {
	GetString(string, ...string) string
}, columns map[int]Data, columnsType []*sql.ColumnType) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Individual column filtering
		columnSearch := ""
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

				column := columns[columnIdx]
				columndb := column.Db

				query := bindingTypes(columndb, str, columnsType)

				if columnSearch != "" && query != "" {
					columnSearch += " AND "
				}

				columnSearch += query

			} else {
				if columnIdx < 0 && requestColumn == "true" {
					fmt.Printf("Do you forgot searchable: false in column %v ?\n", keyColumnsData)
				}
			}
		}
		return db.Where(columnSearch)
	}
}

//Refactor this
func order(c interface {
	GetString(string, ...string) string
}, columns map[int]Data) func(db *gorm.DB) *gorm.DB {
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
					query := fmt.Sprintf("%s %s", column.Db, order)
					db = db.Order(query)
				} else {
					if columnIdx < 0 && c.GetString(columnIdxTittle) == "true" {
						fmt.Printf("Do you forgot orderable: false in column %v ?\n", columnIdxOrder)
					}
				}
			}
		}
		return db
	}
}

func limit(c interface {
	GetString(string, ...string) string
}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		start, err := strconv.Atoi(c.GetString("start"))
		check(err)

		length, err := strconv.Atoi(c.GetString("length"))
		check(err)

		if length < 0 {
			length = 10
		}
		if start < 0 {
			start = 0
		}

		return db.Offset(start).Limit(length)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func search(column map[int]Data, keyColumnsI string) int {
	var i int
	for i = 0; i < len(column); i++ {
		data := column[i]

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
func bindingTypes(columndb string, value string, columnsType []*sql.ColumnType) string {

	for _, element := range columnsType {
		if element.Name() == columndb {

			switch element.ScanType().String() {
			case "string":
				str := "%" + value + "%"
				return fmt.Sprintf("%s LIKE '%s'", columndb, str)
			case "int32":
				intval, err := strconv.Atoi(value)
				check(err)
				return fmt.Sprintf("%s = %d", columndb, intval)
			default:
				fmt.Printf("New type %v\n", element.ScanType().String())
				return ""
			}
		}
	}

	return ""
}

// https://github.com/jinzhu/gorm/issues/1167
func getFields(rows *sql.Rows) map[string]interface{} {

	columns, err := rows.Columns()
	check(err)

	length := len(columns)
	current := makeResultReceiver(length)

	err = rows.Scan(current...)
	check(err)

	value := make(map[string]interface{})
	for i := 0; i < length; i++ {
		key := columns[i]
		val := *(current[i]).(*interface{})
		if val == nil {
			value[key] = nil
			continue
		}
		vType := reflect.TypeOf(val)
		switch vType.String() {
		case "int64":
			value[key] = val.(int64)
		case "string":
			value[key] = val.(string)
		case "time.Time":
			value[key] = val.(time.Time)
		case "[]uint8":
			value[key] = string(val.([]uint8))
		case "bool":
			value[key] = val.(bool)
		default:
			value[key] = val
		}

	}
	return value
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

func initBinding(db *gorm.DB, table string) []*sql.ColumnType {
	rows, err := db.Select("*").
		Table(table).
		Limit(0).
		Rows()
	check(err)

	columnsType, err := rows.ColumnTypes()
	check(err)

	defer rows.Close()
	return columnsType
}
