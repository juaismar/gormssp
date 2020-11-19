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
		},
		{
			Name:       "JuAn",
			Instrument: "Trompeta",
			Age:        15,
			Candies:    -10,
			BirthDate:  date,
			Fun:        false,
		},
		{
			Name:       "Joaquin",
			Instrument: "Flauta",
			Age:        18,
			Candies:    10,
			BirthDate:  date,
			Fun:        true,
		},
		{
			Name:       "Ezequiel",
			Instrument: "Trompeta",
			Age:        13,
			Candies:    5,
			BirthDate:  date,
			Fun:        false,
		},
		{
			Name:       "Marta",
			Instrument: "Tambor",
			Age:        15,
			Candies:    20,
			BirthDate:  date,
			Fun:        false,
		},
		{
			Name:       "Laura",
			Instrument: "Flauta",
			Age:        10,
			Candies:    110,
			BirthDate:  date,
			Fun:        true,
		},
	}
}
