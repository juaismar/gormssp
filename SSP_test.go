package SSP

import (
	"fmt"
	"time"
	"gormssp/test"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "github.com/lib/pq"
)

type Controller struct {
	Params map[string]string
}

func (c *Controller) GetString(key string, def ...string) string{
	return c.Params[key]
}


var _ = Describe("Test for SSP", func() {
	db := initDB()

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

			c := Controller{Params: mapa}

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			result := Simple(&c, db, "users", columns)

			Expect(result.Draw).To(Equal(64))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(2))

			testData := make([]interface{}, 0)
			row := make(map[string]interface{})
			row["0"] = "Juan"
			testData = append(testData, row)
			row = make(map[string]interface{})
			row["0"] = "JuAn"
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
		//filter whereResult (where in only result sended)
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
			whereResult = append(whereResult, "fun IS TRUE")

			whereAll := make([]string, 0)

			result := Complex(&c, db, "users", columns, whereResult, whereAll)

			Expect(result.Draw).To(Equal(62))
			Expect(result.RecordsTotal).To(Equal(6))
			Expect(result.RecordsFiltered).To(Equal(6))

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

	Describe("flated", func() {
		It("returns Empty", func() {

			var whereArray []string

			result := flated(whereArray)

			Expect(result).To(Equal(""))
		})
		It("returns one query", func() {

			var whereArray []string
			whereArray = append(whereArray, "number = 1")

			result := flated(whereArray)

			Expect(result).To(Equal("number = 1"))
		})
		It("returns two query", func() {

			var whereArray []string
			whereArray = append(whereArray, "number = 1")
			whereArray = append(whereArray, "name = 'John'")

			result := flated(whereArray)

			Expect(result).To(Equal("number = 1 AND name = 'John'"))
		})
	})
	Describe("search", func() {
		It("returns -1", func() {

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "role", Dt: "role", Formatter: nil}
			columns[2] = Data{Db: "email", Dt: 2, Formatter: nil}

			result := search(columns, "")

			Expect(result).To(Equal(-1))
		})
		It("returns -1", func() {

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "role", Dt: "role", Formatter: nil}
			columns[2] = Data{Db: "email", Dt: 2, Formatter: nil}

			result := search(columns, "surname")

			Expect(result).To(Equal(-1))
		})
		It("returns 1", func() {

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "role", Dt: "role", Formatter: nil}
			columns[2] = Data{Db: "email", Dt: 2, Formatter: nil}

			result := search(columns, "role")

			Expect(result).To(Equal(1))
		})
		It("returns 0", func() {

			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			columns[1] = Data{Db: "role", Dt: "role", Formatter: nil}
			columns[2] = Data{Db: "email", Dt: 2, Formatter: nil}

			result := search(columns, "0")

			Expect(result).To(Equal(0))
		})

	})
})
func OpenTestConnection() (db *gorm.DB, err error) {
	
	dbDSN := "user=postgress password=postgress DB.name=postgress port=5432 sslmode=disable"
	
	db, err = gorm.Open("postgres", dbDSN)
	
	db.LogMode(true)

	db.DB().SetMaxIdleConns(10)

	return
}

func initDB() *gorm.DB{
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
	  panic(err)
	}
  
	db.AutoMigrate(&model.User{})
  
	db.Unscoped().Delete(&model.User{})
	fillData(db)
	
	return db
}

func fillData(db *gorm.DB){
	db.Create(&model.User{Name: "Juan", Age:10, BirthDate: time.Now(), Fun: true})
	db.Create(&model.User{Name: "JuAn", Age:15, BirthDate: time.Now(), Fun: false})
	db.Create(&model.User{Name: "Joaquin", Age:18, BirthDate: time.Now(), Fun: true})
	db.Create(&model.User{Name: "Ezequiel", Age:13, BirthDate: time.Now(), Fun: false})
	db.Create(&model.User{Name: "Marta", Age:15, BirthDate: time.Now(), Fun: false})
	db.Create(&model.User{Name: "Laura", Age:10, BirthDate: time.Now(), Fun: true})
}