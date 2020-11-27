package databases

import (
	"github.com/jinzhu/gorm"
	model "github.com/juaismar/gormssp/test/models"
)

// InitDB clear and populate database
func InitDB(db *gorm.DB) {

	db.AutoMigrate(&model.User{})

	db.Unscoped().Delete(&model.User{})
	fillData(db)
}

func fillData(db *gorm.DB) {

	for _, user := range model.GetDefaultData() {
		db.Create(&user)
	}

}
