package services

import (
	"blog-server/models"
)

/*
GetAllArticles method
*/
func GetAllArticles() []models.Article {
	db := GetDB()
	articles := make([]models.Article, 0)
	db.Find(&articles).Omit("AuthorID")
	return articles
}

/*
CreateArticle method to create and return the entity of the a article
*/
func CreateArticle(articleToSave *models.Article) *models.Article {
	db := GetDB()
	if db.NewRecord(&articleToSave) {
		// articleToSave.AuthorID = 1
		articleToSave.Author = &models.User{}
		db.Create(&articleToSave)
	}
	return articleToSave
}
