package SSP

import (
	//"fmt"
	"time"
	"gormssp/test"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "github.com/lib/pq"
)

type Controller struct {
	Params map[string]string
	/*draw    string

	start   string
	length  string

	order []interface
	search []interface
	columns []interface*/
}

func (c *Controller) GetString(key string, def ...string) string{
	//value := c.params[key]
	/*if value == nil {
		value = ""
	}*/
	/*if val, ok := c.params[key]; ok {
		
fmt.Printf("val%v\n", val)
		return val
	}
	
fmt.Printf("nada\n")*/
	return ""
}


var _ = Describe("Test for SSP", func() {
	//db := initDB()

	Describe("show name", func() {
		It("returns Empty", func() {
//show name
			var c *Controller
			//c.Params = make(map[string]string)
			//c.Params["val"] = "val"

			Expect(c.Params).To(Equal(false))
			/*
			columns := make(map[int]Data)
			columns[0] = Data{Db: "name", Dt: 0, Formatter: nil}
			//columns[1] = Data{Db: "surname", Dt: 1, Formatter: nil}
			//columns[2] = Data{Db: "age", Dt: 2, Formatter: nil}
			//columns[3] = Data{Db: "birthDate", Dt: 3, Formatter: nil}
			//columns[4] = Data{Db: "fun", Dt: 4, Formatter: nil}

			result := Simple(c, db, "user", columns)

			Expect(result).To(Equal(""))*/
			//Expect(c.GetString("draw")).To(Equal(""))
			/*val, ok := c.Params["draw"]
			Expect(ok).To(Equal(false))
			Expect(val).To(Equal(""))*/

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

	// fmt.Println("testing postgres...")
	//if dbDSN == "" {
		dbDSN := "user=postgress password=postgress DB.name=postgress port=5432 sslmode=disable"
	//}
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
	//defer db.Close()
  
	db.AutoMigrate(&model.User{})
  
	db.Unscoped().Delete(&model.User{})
	fillData(db)
	
	return db
}

func fillData(db *gorm.DB){
	db.Create(&model.User{Name: "Juan", Surname: "Caracola", Age:10, BirthDate: time.Now(), Fun: true})
	db.Create(&model.User{Name: "Juan", Surname: "Tiza", Age:15, BirthDate: time.Now(), Fun: false})
	db.Create(&model.User{Name: "Joaquin", Surname: "Lapiz", Age:18, BirthDate: time.Now(), Fun: true})
	db.Create(&model.User{Name: "Ezequiel", Surname: "Tiza", Age:13, BirthDate: time.Now(), Fun: false})
	db.Create(&model.User{Name: "Marta", Surname: "Clip", Age:15, BirthDate: time.Now(), Fun: false})
	db.Create(&model.User{Name: "Laura", Surname: "Tiza", Age:10, BirthDate: time.Now(), Fun: true})
}