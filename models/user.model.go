package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Blocked   bool
	Confirmed bool
	Email     string
	Username  string
}
