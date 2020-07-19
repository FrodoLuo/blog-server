package services

import (
	"blog-server/models"
)

func CreateComment(commentToSave *models.Comment) {

	db := GetDB()

	if db.NewRecord(&commentToSave) {
		db.Create(&commentToSave)
	}
}
