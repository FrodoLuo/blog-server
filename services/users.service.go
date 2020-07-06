package services

import (
	"blog-server/models"
	"fmt"
)

func GetUserById(id uint) models.User {
	db := GetDB()
	var user models.User
	fmt.Println(user)
	fmt.Println(id)
	db.Where("id = ?", id).First(&user)
	return user
}
