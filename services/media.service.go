package services

import (
	"blog-server/models"
	"io/ioutil"
	"mime/multipart"
)

func GetMediaByTag(tag string, page uint, pageSize uint) []models.Media {
	db := GetDB()
	media := make([]models.Media, 0)
	db.
		Where("tag = ?", tag).
		Offset(page * pageSize).
		Limit(pageSize).
		Find(&media)
	return media
}

func SaveMedia(fileHeader *multipart.FileHeader) models.Media {
	file, _ := fileHeader.Open()
	buffer := make([]byte, fileHeader.Size)
	file.Read(buffer)
	ioutil.WriteFile(fileHeader.Filename, buffer, 0666)
	return models.Media{}
}
