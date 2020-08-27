package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	globalDB "github.com/juaismar/gormssp/test/dbs"
)

func OpenDB() *gorm.DB {

	db, err := gorm.Open("sqlite3", "C:/sqlite/test.db")
	if err != nil {
		panic(err)
	}

	globalDB.InitDB(db)

	return db
}
