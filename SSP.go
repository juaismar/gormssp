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
	Data            []interface{} `json:"data"`
}

func Simple(c interface {
	GetString(string, ...string) string
}, conn *gorm.DB,
	table string,

	columns map[int]Data) MessageDataTable {

	draw := 0
	draw, err := strconv.Atoi(c.GetString("draw"))

	// Build the SQL query string from the request
	rows, err := conn.Select("*").
		Scopes(limit(c), filter(c, columns), order(c, columns)).
		Table(table).
		Rows()

	check(err)

	Datas := dataOutput(columns, rows)

	//search in DDBB recordsFiltered
	var recordsFiltered int
	conn.Scopes(filter(c, columns)).Table(table).Count(&recordsFiltered)

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

	// Build the SQL query string from the request
	whereResultFlated := flated(whereResult)
	whereAllFlated := flated(whereAll)

	rows, err := conn.Select("*").
		Scopes(limit(c), filter(c, columns), order(c, columns)).
		Where(whereResultFlated).
		Where(whereAllFlated).
		Table(table).
		Rows()

	check(err)
	Datas := dataOutput(columns, rows)

	//search in DDBB recordsFiltered
	var recordsFiltered int
	conn.Scopes(filter(c, columns)).
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
	var out []interface{}

	for rows.Next() {
		fields := getFields(rows)

		//row = make(map[interface{}]interface{})
		row := make(map[string]interface{})
		var j = 0
		for j = 0; j < len(columns); j++ {
			column := columns[j]
			var dt string //:= column.Dt

			//TODO? return a interface (int / string)
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
func filter(c interface {
	GetString(string, ...string) string
}, columns map[int]Data) func(db *gorm.DB) *gorm.DB {
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

				column := columns[columnIdx]

				requestColumnQuery := fmt.Sprintf("columns[%d][searchable]", i)
				requestColumn := c.GetString(requestColumnQuery)

				if requestColumn == "true" {
					binding := "%" + str + "%"
					columndb := column.Db

					if globalSearch != "" {
						globalSearch += " OR "
					}
					globalSearch += fmt.Sprintf("%s LIKE '%s'", columndb, binding)
				}
			}
		}
		db = db.Where(globalSearch)

		columnSearch := ""

		// Individual column filtering
		var i int
		for i = 0; ; i++ {
			keyColumnsI := fmt.Sprintf("columns[%d][data]", i)

			keyColumnsData := c.GetString(keyColumnsI)
			if keyColumnsData == "" {
				break
			}

			columnIdx := search(columns, keyColumnsData)

			column := columns[columnIdx]

			requestColumnQuery := fmt.Sprintf("columns[%d][searchable]", i)
			requestColumn := c.GetString(requestColumnQuery)

			requestColumnQuery = fmt.Sprintf("columns[%d][searchable][search][value]", i)
			str := c.GetString(requestColumnQuery)

			if requestColumn == "true" && str != "" {
				binding := "%" + str + "%"
				columndb := column.Db

				if columnSearch != "" {
					columnSearch += " OR "
				}
				columnSearch += fmt.Sprintf("%s LIKE '%s'", columndb, binding)
			}
		}
		return db.Where(columnSearch)
	}
}

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

				column := columns[columnIdx]

				columnIdxTittle = fmt.Sprintf("columns[%s][orderable]", columnIdxOrder)
				if c.GetString(columnIdxTittle) == "true" {

					columnIdxTittle = fmt.Sprintf("order[%d][dir]", i)
					requestColumnData = c.GetString(columnIdxTittle)

					order := "desc"
					if requestColumnData == "asc" {
						order = "asc"
					}

					query := fmt.Sprintf("%s %s", column.Db, order)
					db = db.Order(query)
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
