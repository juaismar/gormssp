package SSP

//This file is comented until configure travish sqlite3
import (
	"fmt"
	"time"

	sqlite "github.com/juaismar/gormssp/test/dbs/sqlite"

	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test for SSP", func() {
	db := sqlite.OpenDB()

	Describe("Simple", func() {
		It("returns from 0 to 4", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "62"
			mapa["start"] = "0"
			mapa["length"] = "4"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			result := Simple(&c, db, "users", columns)

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
		//length is negative
		It("returns from 10 elements", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "62"
			mapa["start"] = "0"
			mapa["length"] = "-1"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			result := Simple(&c, db, "users", columns)

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
		//start is negative
		It("returns from 0 to 4", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "62"
			mapa["start"] = "-1"
			mapa["length"] = "4"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			result := Simple(&c, db, "users", columns)

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
		It("returns from 2 to 6", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "63"
			mapa["start"] = "2"
			mapa["length"] = "4"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			result := Simple(&c, db, "users", columns)

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
		//global search
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

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "instrument", Dt: 1, Formatter: nil}
			columns[2] = Data{Db: "age", Dt: 2, Formatter: nil}
			result := Simple(&c, db, "users", columns)

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
		//naming a row
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

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: "supername", Formatter: nil}
			result := Simple(&c, db, "users", columns)

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
		//search LIKE string case insensitive
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

			mapa["columns[1][data]"] = "0"
			mapa["columns[1][searchable]"] = "true"
			mapa["columns[1][search][value]"] = ""

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "instrument", Dt: 1, Formatter: nil}
			result := Simple(&c, db, "users", columns)

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
		//search LIKE string case sensitive
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

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Cs: true, Formatter: nil}
			result := Simple(&c, db, "users", columns)

			Expect(result.Draw).To(Equal(64))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(1))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = "JuAn"
			testData = append(testData, row)

			Expect(result.Data).To(Equal(testData))
		})

		//search int
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

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "age", Dt: 1, Formatter: nil}
			result := Simple(&c, db, "users", columns)

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
		//search bool
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

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "fun", Dt: 1, Formatter: nil}
			result := Simple(&c, db, "users", columns)

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
		//test format
		It("return name whit prefix and age", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "62"
			mapa["start"] = "0"
			mapa["length"] = "4"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: func(
				data interface{}, row map[string]interface{}) interface{} {
				return fmt.Sprintf("PREFIX_%v_%v", data, row["age"])
			}}

			result := Simple(&c, db, "users", columns)

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

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "instrument", Dt: 1, Formatter: nil}
			result := Simple(&c, db, "users", columns)

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

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "instrument", Dt: 1, Formatter: nil}
			result := Simple(&c, db, "users", columns)

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
		//TODO uncoment when compatible
		/*XIt("Global search regex", func() {

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

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "instrument", Dt: 1, Formatter: nil}
			result := Simple(&c, db, "users", columns)

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
		XIt("returns names whit 5 chars (regex)", func() {

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

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			result := Simple(&c, db, "users", columns)

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
		})*/
	})
	Describe("Complex", func() {
		//filter whereall (where in all queries)
		It("returns fun only Juan Joaquin Laura", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "62"
			mapa["start"] = "0"
			mapa["length"] = "4"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}

			whereResult := make([]string, 0)

			whereAll := make([]string, 0)
			whereAll = append(whereAll, "fun IS TRUE")

			result := Complex(&c, db, "users", columns, whereResult, whereAll)

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

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}

			whereResult := make([]string, 0)
			whereResult = append(whereResult, "fun IS TRUE")

			whereAll := make([]string, 0)

			result := Complex(&c, db, "users", columns, whereResult, whereAll)

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

		It("returns a time", func() {

			mapa := make(map[string]string)
			mapa["draw"] = "62"
			mapa["start"] = "0"
			mapa["length"] = "1"
			mapa["order[0][column]"] = "0"
			mapa["order[0][dir]"] = "asc"

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "birth_date", Dt: 0, Formatter: func(
				data interface{}, row map[string]interface{}) interface{} {

				layoutISO := "2006-01-02"
				testTime, _ := time.Parse(layoutISO, "2011-11-11")

				time := data.(time.Time)

				Expect(time.Equal(testTime)).To(BeTrue())
				return data
			}}
			result := Simple(&c, db, "users", columns)

			Expect(result.Draw).To(Equal(62))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(6))
		})
	})
})
