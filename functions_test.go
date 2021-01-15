package ssp_test

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	ssp "github.com/juaismar/gormssp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// ControllerEmulated emulate the beego controller
type ControllerEmulated struct {
	Params map[string]string
}

// GetString emulate the beego controoller method
func (c *ControllerEmulated) GetString(key string, def ...string) string {
	return c.Params[key]
}

// FunctionsTest internal function test
func FunctionsTest() {
	Describe("flated", func() {
		It("returns Empty", func() {

			var whereArray []string

			result := ssp.Flated(whereArray)

			Expect(result).To(Equal(""))
		})
		It("returns one query", func() {

			var whereArray []string
			whereArray = append(whereArray, "number = 1")

			result := ssp.Flated(whereArray)

			Expect(result).To(Equal("number = 1"))
		})
		It("returns two query", func() {

			var whereArray []string
			whereArray = append(whereArray, "number = 1")
			whereArray = append(whereArray, "name = 'John'")

			result := ssp.Flated(whereArray)

			Expect(result).To(Equal("number = 1 AND name = 'John'"))
		})
	})
	Describe("search", func() {
		It("returns -1", func() {

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = ssp.Data{Db: "role", Dt: "role", Formatter: nil}
			columns[2] = ssp.Data{Db: "email", Dt: 2, Formatter: nil}

			result := ssp.Search(columns, "")

			Expect(result).To(Equal(-1))
		})
		It("returns -1", func() {

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = ssp.Data{Db: "role", Dt: "role", Formatter: nil}
			columns[2] = ssp.Data{Db: "email", Dt: 2, Formatter: nil}

			result := ssp.Search(columns, "instrument")

			Expect(result).To(Equal(-1))
		})
		It("returns 1", func() {

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = ssp.Data{Db: "role", Dt: "role", Formatter: nil}
			columns[2] = ssp.Data{Db: "email", Dt: 2, Formatter: nil}

			result := ssp.Search(columns, "role")

			Expect(result).To(Equal(1))
		})
		It("returns 0", func() {

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = ssp.Data{Db: "role", Dt: "role", Formatter: nil}
			columns[2] = ssp.Data{Db: "email", Dt: 2, Formatter: nil}

			result := ssp.Search(columns, "0")

			Expect(result).To(Equal(0))
		})

	})
}

// ComplexFunctionTest test for Complex method
func ComplexFunctionTest(db *gorm.DB) {
	Describe("Complex", func() {
		//filter whereall (where in all queries)
		It("returns fun only Juan Joaquin Laura", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "62"
			mapa["start"] = "0"
			mapa["length"] = "4"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}

			whereResult := make([]string, 0)

			whereAll := make([]string, 0)
			whereAll = append(whereAll, "fun IS TRUE")

			result, err := ssp.Complex(&c, db, "users", columns, whereResult, whereAll)

			Expect(err).To(BeNil())
			Expect(result.Draw).To(Equal(62))
			Expect(result.RecordsTotal).To(Equal(3))
			Expect(result.RecordsFiltered).To(Equal(3))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = "Juan"
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = "Joaquin"
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = "Laura"
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})
		//filter whereResult (where in only result sended)
		It("returns fun only Juan Joaquin Laura", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "62"
			mapa["start"] = "0"
			mapa["length"] = "5"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}

			whereResult := make([]string, 0)
			whereResult = append(whereResult, "fun IS TRUE")

			whereAll := make([]string, 0)

			result, err := ssp.Complex(&c, db, "users", columns, whereResult, whereAll)

			Expect(err).To(BeNil())
			Expect(result.Draw).To(Equal(62))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(3))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = "Juan"
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = "Joaquin"
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = "Laura"
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})
	})
}

// RegExpTest test for regular expression
func RegExpTest(db *gorm.DB) {
	Describe("RegExp", func() {
		It("Global search regex", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "64"
			mapa["start"] = "0"
			mapa["length"] = "10"
			mapa["order[0][column]"] = "1"
			mapa["order[0][dir]"] = "desc"

			mapa["search[value]"] = "^Eze"
			mapa["search[regex]"] = "true"

			mapa["columns[0][data]"] = "0"
			mapa["columns[0][searchable]"] = "true"

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = ssp.Data{Db: "instrument", Dt: 1, Formatter: nil}
			result, err := ssp.Simple(&c, db, "users", columns)

			Expect(err).To(BeNil())
			Expect(result.Draw).To(Equal(64))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(1))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = "Ezequiel"
			row["1"] = "Trompeta"
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})
		It("returns names whit 5 chars (regex)", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "64"
			mapa["start"] = "0"
			mapa["length"] = "10"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			mapa["columns[0][data]"] = "0"
			mapa["columns[0][searchable]"] = "true"
			mapa["columns[0][search][value]"] = "^.{5}$"
			mapa["columns[0][search][regex]"] = "true"

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
			result, err := ssp.Simple(&c, db, "users", columns)

			Expect(err).To(BeNil())
			Expect(result.Draw).To(Equal(64))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(2))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = "Marta"
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = "Laura"
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})
		It("returns names 2 names", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "64"
			mapa["start"] = "0"
			mapa["length"] = "10"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			mapa["columns[0][data]"] = "0"
			mapa["columns[0][searchable]"] = "true"
			mapa["columns[0][search][value]"] = "Marta|Laura"
			mapa["columns[0][search][regex]"] = "true"

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
			result, err := ssp.Simple(&c, db, "users", columns)

			Expect(err).To(BeNil())
			Expect(result.Draw).To(Equal(64))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(2))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = "Marta"
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = "Laura"
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})
		It("returns 2 ages int", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "64"
			mapa["start"] = "0"
			mapa["length"] = "10"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			mapa["columns[0][data]"] = "0"
			mapa["columns[0][searchable]"] = "true"
			mapa["columns[0][search][value]"] = "13|18"
			mapa["columns[0][search][regex]"] = "true"

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "age", Dt: 0, Formatter: nil}
			result, err := ssp.Simple(&c, db, "users", columns)

			Expect(err).To(BeNil())
			Expect(result.Draw).To(Equal(64))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(2))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = int64(18)
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = int64(13)
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})
		FIt("returns 2 money float", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "64"
			mapa["start"] = "0"
			mapa["length"] = "10"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			mapa["columns[0][data]"] = "0"
			mapa["columns[0][searchable]"] = "true"
			mapa["columns[0][search][value]"] = "22.1100|0.1000"
			mapa["columns[0][search][regex]"] = "true"

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "money", Dt: 0, Formatter: nil}
			result, err := ssp.Simple(&c, db, "users", columns)

			Expect(err).To(BeNil())
			Expect(result.Draw).To(Equal(64))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(2))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = float64(22.110000610351562)
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = float64(0.10000000149011612)
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})
	})
}

// Types test for types
func Types(db *gorm.DB) {
	Describe("Types", func() {
		Describe("uint", func() {
			It("returns 2 Age 15", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = ""

				mapa["columns[1][data]"] = "1"
				mapa["columns[1][searchable]"] = "true"
				mapa["columns[1][search][value]"] = "15"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "age", Dt: 1, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(2))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "JuAn"
				row["1"] = int64(15)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Marta"
				row["1"] = int64(15)
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("int", func() {
			It("returns 1 Candies 10", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = ""

				mapa["columns[1][data]"] = "1"
				mapa["columns[1][searchable]"] = "true"
				mapa["columns[1][search][value]"] = "10"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "candies", Dt: 1, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(1))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Joaquin"
				row["1"] = int64(10)
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("int 8", func() {
			It("returns 2 users", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = ""

				mapa["columns[1][data]"] = "1"
				mapa["columns[1][searchable]"] = "true"
				mapa["columns[1][search][value]"] = "1"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "toys", Dt: 1, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(2))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "JuAn"
				row["1"] = int64(1)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Marta"
				row["1"] = int64(1)
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("bool", func() {
			It("returns fun only Juan Joaquin Laura", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = ""

				mapa["columns[1][data]"] = "1"
				mapa["columns[1][searchable]"] = "true"
				mapa["columns[1][search][value]"] = "true"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "fun", Dt: 1, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(3))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Juan"
				row["1"] = true
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Joaquin"
				row["1"] = true
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Laura"
				row["1"] = true
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("float32", func() {
			It("returns money only Juan Marta", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = ""

				mapa["columns[1][data]"] = "1"
				mapa["columns[1][searchable]"] = "true"
				mapa["columns[1][search][value]"] = "2.0"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "money", Dt: 1, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(2))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Juan"
				row["1"] = float64(2.0)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Marta"
				row["1"] = float64(2.0)
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
			It("returns all with decimals", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = ""

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "money", Dt: 1, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(6))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Juan"
				row["1"] = float64(2.0)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "JuAn"
				row["1"] = float64(3.0999999046325684)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Joaquin"
				row["1"] = float64(3.4000000953674316)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Ezequiel"
				row["1"] = float64(22.110000610351562)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Marta"
				row["1"] = float64(2.0)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Laura"
				row["1"] = float64(0.10000000149011612)
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("float64", func() {
			It("returns bitcoins only Juan Marta", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = ""

				mapa["columns[1][data]"] = "1"
				mapa["columns[1][searchable]"] = "true"
				mapa["columns[1][search][value]"] = "3.0"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "bitcoins", Dt: 1, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(2))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Juan"
				row["1"] = float64(3.0)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Marta"
				row["1"] = float64(3.0)
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
			It("returns all with decimals", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = ""

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "bitcoins", Dt: 1, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(6))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Juan"
				row["1"] = float64(3.0)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "JuAn"
				row["1"] = float64(4.3)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Joaquin"
				row["1"] = float64(7.18)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Ezequiel"
				row["1"] = float64(82.14)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Marta"
				row["1"] = float64(3.0)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Laura"
				row["1"] = float64(22.71)
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("time.TIME", func() {
			It("returns a time and formatter", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "62"
				mapa["start"] = "0"
				mapa["length"] = "1"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "birth_date", Dt: 0, Formatter: func(
					data interface{}, row map[string]interface{}) (interface{}, error) {

					layoutISO := "2006-01-02"
					testTime, err := time.Parse(layoutISO, "2011-11-11")

					time := data.(time.Time)

					Expect(time.Equal(testTime)).To(BeTrue())
					return time, err
				}}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(62))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(6))
			})
		})
		Describe("UUID", func() {
			It("returns Juan", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = ""

				mapa["columns[1][data]"] = "1"
				mapa["columns[1][searchable]"] = "true"
				mapa["columns[1][search][value]"] = "bfe44cb2-c65c-4f37-9672-8437b6718d70"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "uuid", Dt: 1, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(1))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Juan"
				row["1"] = "bfe44cb2-c65c-4f37-9672-8437b6718d70"
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
	})
}

// SimpleFunctionTest test for ssp.Simplex method
func SimpleFunctionTest(db *gorm.DB) {
	Describe("Simple and basic features", func() {
		It("returns from 0 to 4", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "62"
			mapa["start"] = "0"
			mapa["length"] = "4"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
			result, err := ssp.Simple(&c, db, "users", columns)

			Expect(err).To(BeNil())
			Expect(result.Draw).To(Equal(62))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(6))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = "Juan"
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = "JuAn"
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = "Joaquin"
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = "Ezequiel"
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})
		Describe("Length is negative", func() {
			It("returns from 10 elements", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "62"
				mapa["start"] = "0"
				mapa["length"] = "-1"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(62))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(6))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Juan"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "JuAn"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Joaquin"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Ezequiel"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Marta"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Laura"
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("Start is negative", func() {
			It("returns from 0 to 4", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "62"
				mapa["start"] = "-1"
				mapa["length"] = "4"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(62))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(6))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Juan"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "JuAn"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Joaquin"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Ezequiel"
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("Paginate", func() {
			It("returns from 2 to 6", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "63"
				mapa["start"] = "2"
				mapa["length"] = "4"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(63))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(6))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Joaquin"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Ezequiel"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Marta"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Laura"
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("Global search", func() {
			It("returns 2 Juan", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["search[value]"] = "uAn"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = ""

				mapa["columns[1][data]"] = "1"
				mapa["columns[1][searchable]"] = "true"
				mapa["columns[1][search][value]"] = ""

				mapa["columns[2][data]"] = "2"
				mapa["columns[2][searchable]"] = "true"
				mapa["columns[2][search][value]"] = ""

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "instrument", Dt: 1, Formatter: nil}
				columns[2] = ssp.Data{Db: "age", Dt: 2, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(2))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Juan"
				row["1"] = "Tambor"
				row["2"] = int64(10)
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "JuAn"
				row["1"] = "Trompeta"
				row["2"] = int64(15)
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("Multiple individual search", func() {
			It("returns 1 Juan", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = "Juan"

				mapa["columns[1][data]"] = "1"
				mapa["columns[1][searchable]"] = "true"
				mapa["columns[1][search][value]"] = "Tambor"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "instrument", Dt: 1, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(1))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Juan"
				row["1"] = "Tambor"
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("Naming a row", func() {
			It("returns all", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "3"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[supername][data]"] = "0"
				mapa["columns[supername][searchable]"] = "true"
				mapa["columns[supername][search][value]"] = ""

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: "supername", Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(6))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["supername"] = "Juan"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["supername"] = "JuAn"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["supername"] = "Joaquin"
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("Search LIKE string case insensitive", func() {
			It("returns 2 Juan", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = "uAn"

				mapa["columns[1][data]"] = "1"
				mapa["columns[1][searchable]"] = "true"
				mapa["columns[1][search][value]"] = ""

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "instrument", Dt: 1, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(2))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Juan"
				row["1"] = "Tambor"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "JuAn"
				row["1"] = "Trompeta"
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("Search on varchar LIKE string case insensitive", func() {
			It("returns 2 Tambor", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "1"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = "ambor"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
				columns[1] = ssp.Data{Db: "instrument", Dt: 1, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(2))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "Juan"
				row["1"] = "Tambor"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "Marta"
				row["1"] = "Tambor"
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("Search LIKE string case sensitive", func() {
			It("returns 2 Juan", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "64"
				mapa["start"] = "0"
				mapa["length"] = "10"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				mapa["columns[0][data]"] = "0"
				mapa["columns[0][searchable]"] = "true"
				mapa["columns[0][search][value]"] = "uAn"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Cs: true, Formatter: nil}
				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(64))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(1))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "JuAn"
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		Describe("Format", func() {
			It("return name whit prefix and age", func() {

				mapa := make(map[string]string)
				mapa["draw"] = "62"
				mapa["start"] = "0"
				mapa["length"] = "4"
				mapa["order[0][column]"] = "0"
				mapa["order[0][dir]"] = "asc"

				c := ControllerEmulated{Params: mapa}

				columns := make(map[int]ssp.Data)
				columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: func(
					data interface{}, row map[string]interface{}) (interface{}, error) {
					return fmt.Sprintf("PREFIX_%v_%v", data, row["age"]), nil
				}}

				result, err := ssp.Simple(&c, db, "users", columns)

				Expect(err).To(BeNil())
				Expect(result.Draw).To(Equal(62))
				Expect(result.RecordsTotal).To(Equal(6))
				Expect(result.RecordsFiltered).To(Equal(6))

				testData := make([]interface{}, 0)
				row := make(map[string]interface{})
				row["0"] = "PREFIX_Juan_10"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "PREFIX_JuAn_15"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "PREFIX_Joaquin_18"
				testData = append(testData, row)
				row = make(map[string]interface{})
				row["0"] = "PREFIX_Ezequiel_13"
				testData = append(testData, row)

				Expect(result.Data).To(Equal(testData))
			})
		})
		It("Ordering by instrument asc", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "64"
			mapa["start"] = "0"
			mapa["length"] = "10"
			mapa["order[0][column]"] = "1"
			mapa["order[0][dir]"] = "asc"

			mapa["search[value]"] = "uAn"

			mapa["columns[0][data]"] = "0"
			mapa["columns[0][searchable]"] = "true"
			mapa["columns[0][orderable]"] = "true"
			mapa["columns[0][search][value]"] = ""

			mapa["columns[1][data]"] = "0"
			mapa["columns[1][searchable]"] = "true"
			mapa["columns[1][orderable]"] = "true"
			mapa["columns[1][search][value]"] = ""

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = ssp.Data{Db: "instrument", Dt: 1, Formatter: nil}
			result, err := ssp.Simple(&c, db, "users", columns)

			Expect(err).To(BeNil())
			Expect(result.Draw).To(Equal(64))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(2))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = "Juan"
			row["1"] = "Tambor"
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = "JuAn"
			row["1"] = "Trompeta"
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})
		It("Ordering by instrument desc", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "64"
			mapa["start"] = "0"
			mapa["length"] = "10"
			mapa["order[0][column]"] = "1"
			mapa["order[0][dir]"] = "desc"

			mapa["search[value]"] = "uAn"

			mapa["columns[0][data]"] = "0"
			mapa["columns[0][searchable]"] = "true"
			mapa["columns[0][orderable]"] = "true"
			mapa["columns[0][search][value]"] = ""

			mapa["columns[1][data]"] = "0"
			mapa["columns[1][searchable]"] = "true"
			mapa["columns[1][orderable]"] = "true"
			mapa["columns[1][search][value]"] = ""

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = ssp.Data{Db: "instrument", Dt: 1, Formatter: nil}
			result, err := ssp.Simple(&c, db, "users", columns)

			Expect(err).To(BeNil())
			Expect(result.Draw).To(Equal(64))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(2))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = "JuAn"
			row["1"] = "Trompeta"
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = "Juan"
			row["1"] = "Tambor"
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})
	})
}

// Errors test some errors
func Errors(db *gorm.DB) {
	Describe("Column not found", func() {
		It("Return error", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "64"
			mapa["start"] = "0"
			mapa["length"] = "2"
			mapa["order[0][column]"] = "1"
			mapa["order[0][dir]"] = "desc"

			mapa["columns[0][data]"] = "0"
			mapa["columns[0][searchable]"] = "true"

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "bike", Dt: 0, Formatter: nil}
			result, err := ssp.Simple(&c, db, "users", columns)

			Expect(err).To(BeNil())

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = nil
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = nil
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})
	})
	Describe("Bad map id", func() {
		It("Return error", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "64"
			mapa["start"] = "0"
			mapa["length"] = "2"
			mapa["order[0][column]"] = "1"
			mapa["order[0][dir]"] = "desc"

			mapa["columns[0][data]"] = "0"
			mapa["columns[0][searchable]"] = "true"

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[1] = ssp.Data{Db: "bike", Dt: 0, Formatter: nil}
			_, err := ssp.Simple(&c, db, "users", columns)

			Expect(fmt.Sprintf("%v", err)).To(Equal("Bad map id, column[0] dont exist"))
		})
	})
	Describe("Format error", func() {
		It("return name whit prefix and age", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "62"
			mapa["start"] = "0"
			mapa["length"] = "4"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: func(
				data interface{}, row map[string]interface{}) (interface{}, error) {
				layout := "2006-01-02T15:04:05.000Z"
				//try convert name to date
				return time.Parse(layout, data.(string))
			}}

			_, err := ssp.Simple(&c, db, "users", columns)

			Expect(err).ToNot(BeNil())
		})
	})
	Describe("Column with reserved word", func() {
		It("returns 2 Age 15", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "64"
			mapa["start"] = "0"
			mapa["length"] = "10"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			mapa["columns[0][data]"] = "0"
			mapa["columns[0][searchable]"] = "true"
			mapa["columns[0][search][value]"] = ""

			mapa["columns[1][data]"] = "1"
			mapa["columns[1][searchable]"] = "true"
			mapa["columns[1][search][value]"] = "2"

			c := ControllerEmulated{Params: mapa}

			columns := make(map[int]ssp.Data)
			columns[0] = ssp.Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = ssp.Data{Db: "end", Dt: 1, Formatter: nil}
			result, err := ssp.Simple(&c, db, "users", columns)

			Expect(err).To(BeNil())
			Expect(result.Draw).To(Equal(64))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(1))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = "Joaquin"
			row["1"] = int64(2)
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})
	})
}
