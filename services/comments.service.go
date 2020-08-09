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

func GetComments(page uint, pageSize uint) []models.Comment {
	comments := make([]models.Comment, 0)
	GetDB().
		Preload("Article").
		Offset(page * pageSize).
		Limit(pageSize).
		Find(&comments)

	return comments
}

func CountComments() uint {
	var count uint
	GetDB().
		Model(&models.Comment{}).
		Count(&count)

	return count
}
