package databases

import (
	"github.com/jinzhu/gorm"
	model "github.com/juaismar/gormssp/test/models"
)

// InitDB clear and populate database
func InitDB(db *gorm.DB) {

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Pet{})

	db.Unscoped().Delete(&model.User{})
	db.Unscoped().Delete(&model.Pet{})
	fillData(db)
}

func fillData(db *gorm.DB) {

	for _, user := range model.GetDefaultUser() {
		db.Create(&user)
	}
	for _, pet := range model.GetDefaultPet() {
		db.Create(&pet)
	}

}
