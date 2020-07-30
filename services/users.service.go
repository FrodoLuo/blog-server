package services

import (
	"blog-server/models"
	"crypto/md5"
	"encoding/hex"
	"time"
)

func GetUserById(id uint) models.User {
	db := GetDB()
	user := models.User{}
	db.Where("id = ?", id).First(&user)
	return user
}

func GetUserByToken(token string) models.User {
	db := GetDB()
	user := models.User{}

	db.Where("token = ?", token).First(&user)
	return user
}

func UpdateAuth(email string, password string, key string) (token string, status int) {
	db := GetDB()
	user := models.User{}

	decipheredPass, err := Decipher(password, key)

	if err != nil {
		return "", 401
	}

	db.
		Where("email = ?", email).
		Where("password = ?", decipheredPass).
		First(&user)

	if user.ID == 0 {
		return "", 403
	}
	token = generateToken(email)
	user.Token = token
	db.
		Save(user)
	return token, 0
}

func generateToken(email string) string {
	hash := md5.New()
	hash.Write([]byte(email))
	hash.Write([]byte(string(time.Now().Nanosecond())))
	return hex.EncodeToString(hash.Sum(nil))
}
