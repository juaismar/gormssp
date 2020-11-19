package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

var layoutISO = "2006-01-02"

type User struct {
	gorm.Model

	Name       string
	Instrument string `gorm:"type:varchar(16)"`
	Age        uint
	Candies    int
	BirthDate  time.Time
	Fun        bool
	Money      float32
	Bitcoins   float64
}

func GetDefaultData() []User {

	date, _ := time.Parse(layoutISO, "2011-11-11")
	return []User{
		{
			Name:       "Juan",
			Instrument: "Tambor",
			Age:        10,
			Candies:    0,
			BirthDate:  date,
			Fun:        true,
			Money:      2.0,
			Bitcoins:   3.0,
		},
		{
			Name:       "JuAn",
			Instrument: "Trompeta",
			Age:        15,
			Candies:    -10,
			BirthDate:  date,
			Fun:        false,
			Money:      3.1,
			Bitcoins:   4.3,
		},
		{
			Name:       "Joaquin",
			Instrument: "Flauta",
			Age:        18,
			Candies:    10,
			BirthDate:  date,
			Fun:        true,
			Money:      3.4,
			Bitcoins:   7.18,
		},
		{
			Name:       "Ezequiel",
			Instrument: "Trompeta",
			Age:        13,
			Candies:    5,
			BirthDate:  date,
			Fun:        false,
			Money:      22.11,
			Bitcoins:   82.14,
		},
		{
			Name:       "Marta",
			Instrument: "Tambor",
			Age:        15,
			Candies:    20,
			BirthDate:  date,
			Fun:        false,
			Money:      2.0,
			Bitcoins:   3.0,
		},
		{
			Name:       "Laura",
			Instrument: "Flauta",
			Age:        10,
			Candies:    110,
			BirthDate:  date,
			Fun:        true,
			Money:      0.1,
			Bitcoins:   22.71,
		},
	}
}
