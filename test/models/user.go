package model

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

var layoutISO = "2006-01-02"

// User is the model for test
type User struct {
	gorm.Model

	UUID         uuid.UUID
	Name         string
	Instrument   string `gorm:"type:varchar(16)"`
	Age          uint
	Candies      int
	BirthDate    time.Time
	Fun          bool
	Money        float32
	Bitcoins     float64
	Toys         int64
	End          int
	Secret       []byte
	FavoriteSong string `gorm:"column:Favorite song"`
}

// GetDefaultUser returns data to populate table
func GetDefaultUser() []User {

	date, _ := time.Parse(layoutISO, "2011-11-11")

	uuidJuan, _ := uuid.FromString("bfe44cb2-c65c-4f37-9672-8437b6718d70")
	uuidJuAn, _ := uuid.FromString("c14be350-6671-4ffe-8108-608ebcccf036")
	uuidJoaquin, _ := uuid.FromString("66d13290-ef29-47f8-a291-5bb6474bcc78")
	uuidEzequiel, _ := uuid.FromString("d7ee5bc2-d112-424c-b213-f3d4bc5989ef")
	uuidMarta, _ := uuid.FromString("d1adfebc-8048-4db0-9b9b-c03f3eb5a9d4")
	uuidLaura, _ := uuid.FromString("e4e1f721-c13e-4b7e-a711-887f31570a74")

	return []User{
		{
			UUID:         uuidJuan,
			Name:         "Juan",
			Instrument:   "Tambor",
			Age:          10,
			Candies:      0,
			BirthDate:    date,
			Fun:          true,
			Money:        2.0,
			Bitcoins:     3.0,
			Toys:         0,
			End:          0,
			FavoriteSong: "Happy birthday",
		},
		{
			UUID:         uuidJuAn,
			Name:         "JuAn",
			Instrument:   "Trompeta",
			Age:          15,
			Candies:      -10,
			BirthDate:    date,
			Fun:          false,
			Money:        3.1,
			Bitcoins:     4.3,
			Toys:         1,
			End:          1,
			FavoriteSong: "Himno Español",
		},
		{
			UUID:         uuidJoaquin,
			Name:         "Joaquin",
			Instrument:   "Flauta",
			Age:          18,
			Candies:      10,
			BirthDate:    date,
			Fun:          true,
			Money:        3.4,
			Bitcoins:     7.18,
			Toys:         2,
			End:          2,
			FavoriteSong: "гимн России",
		},
		{
			UUID:         uuidEzequiel,
			Name:         "Ezequiel",
			Instrument:   "Trompeta",
			Age:          13,
			Candies:      5,
			BirthDate:    date,
			Fun:          false,
			Money:        22.11,
			Bitcoins:     82.14,
			Toys:         2,
			End:          3,
			FavoriteSong: "Viejoven",
		},
		{
			UUID:         uuidMarta,
			Name:         "Marta",
			Instrument:   "Tambor",
			Age:          15,
			Candies:      20,
			BirthDate:    date,
			Fun:          false,
			Money:        2.0,
			Bitcoins:     3.0,
			Toys:         1,
			End:          4,
			FavoriteSong: "El chiki chiki",
		},
		{
			UUID:         uuidLaura,
			Name:         "Laura",
			Instrument:   "Flauta",
			Age:          10,
			Candies:      110,
			BirthDate:    date,
			Fun:          true,
			Money:        0.1,
			Bitcoins:     22.71,
			Toys:         3,
			End:          5,
			FavoriteSong: "Street Fighter",
		},
	}
}
