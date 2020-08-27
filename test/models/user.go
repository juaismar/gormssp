package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

var layoutISO = "2006-01-02"

type User struct {
	gorm.Model

	Name       string
	Instrument string
	Age        int
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
			BirthDate:  date,
			Fun:        true,
		},
		{
			Name:       "JuAn",
			Instrument: "Trompeta",
			Age:        15,
			BirthDate:  date,
			Fun:        false,
		},
		{
			Name:       "Joaquin",
			Instrument: "Flauta",
			Age:        18,
			BirthDate:  date,
			Fun:        true,
		},
		{
			Name:       "Ezequiel",
			Instrument: "Trompeta",
			Age:        13,
			BirthDate:  date,
			Fun:        false,
		},
		{
			Name:       "Marta",
			Instrument: "Tambor",
			Age:        15,
			BirthDate:  date,
			Fun:        false,
		},
		{
			Name:       "Laura",
			Instrument: "Flauta",
			Age:        10,
			BirthDate:  date,
			Fun:        true,
		},
	}
}
