package postgres

import (
	"github.com/jinzhu/gorm"

	databases "github.com/juaismar/gormssp/v2/test/dbs"
)

// OpenDB return the Database connection
func OpenDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	databases.InitDB(db)

	return db
}
