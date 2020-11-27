package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Needed for Gorm

	databases "github.com/juaismar/gormssp/test/dbs"
)

// OpenDB return the Database connection
func OpenDB() *gorm.DB {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	databases.InitDB(db)

	return db
}
