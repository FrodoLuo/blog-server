package services

import (
	"blog-server/models"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"os/user"
	"path"
	"path/filepath"
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

func SaveMedia(fileHeader *multipart.FileHeader, tag string, description string, order uint) models.Media {
	file, _ := fileHeader.Open()
	buffer := make([]byte, fileHeader.Size)
	file.Read(buffer)

	dataFolder := initDataFolder()

	err := ioutil.WriteFile(filepath.Join(dataFolder, fileHeader.Filename), buffer, 0666)

	if err != nil {
		fmt.Println(err)
	}

	media := models.Media{
		Source:         path.Join("/files", fileHeader.Filename),
		Tag:            tag,
		Description:    description,
		OrderReference: order,
	}

	GetDB().
		Save(&media)

	return media
}

func initDataFolder() string {
	currentInfo, err := user.Current()
	if err != nil {
		panic(err)
	}
	dataFolder := filepath.Join(currentInfo.HomeDir, "data")

	if !IsPathExist(dataFolder) {
		os.MkdirAll(dataFolder, os.ModePerm)
	}
	return dataFolder
}
