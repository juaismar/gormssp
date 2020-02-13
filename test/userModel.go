package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name      string
	Surname   string
	Age		  int
	BirthDate time.Time
	Fun		  bool
}

