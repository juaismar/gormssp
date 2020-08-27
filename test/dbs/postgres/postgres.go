package postgres

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	globalDB "github.com/juaismar/gormssp/test/dbs"
)

func OpenDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	globalDB.InitDB(db)

	return db
}
