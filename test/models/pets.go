package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Pet is the model for test
type Pet struct {
	gorm.Model

	MasterID uuid.UUID
	Name     string
}

// GetDefaultPet returns data to populate table
func GetDefaultPet() []Pet {

	uuidJuan, _ := uuid.FromString("bfe44cb2-c65c-4f37-9672-8437b6718d70")
	uuidJuAn, _ := uuid.FromString("c14be350-6671-4ffe-8108-608ebcccf036")
	uuidJoaquin, _ := uuid.FromString("66d13290-ef29-47f8-a291-5bb6474bcc78")
	uuidEzequiel, _ := uuid.FromString("d7ee5bc2-d112-424c-b213-f3d4bc5989ef")
	uuidMarta, _ := uuid.FromString("d1adfebc-8048-4db0-9b9b-c03f3eb5a9d4")
	uuidLaura, _ := uuid.FromString("e4e1f721-c13e-4b7e-a711-887f31570a74")

	return []Pet{
		{
			MasterID: uuidJuan,
			Name:     "Cerverus",
		},
		{
			MasterID: uuidJuAn,
			Name:     "Mikey",
		},
		{
			MasterID: uuidJoaquin,
			Name:     "Epona",
		},
		{
			MasterID: uuidEzequiel,
			Name:     "Shadowfax",
		},
		{
			MasterID: uuidMarta,
			Name:     "Rocinante",
		},
		{
			MasterID: uuidLaura,
			Name:     "Tweety",
		},
	}
}
