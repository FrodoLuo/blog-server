package services

import (
	"blog-server/models"
)

/**
GetAllAricles method return all articles of the blog
*/
func GetAllArticles() []models.Article {
	db := GetDB()
	articles := make([]models.Article, 0)
	db.Find(&articles)
	return articles
}
