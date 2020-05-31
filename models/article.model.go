package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model

	Brief   string
	Class   string
	Content string
	Tags    string
	Title   string

	Author User
}
