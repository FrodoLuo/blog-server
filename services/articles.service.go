package services

import (
	"blog-server/models"
	"fmt"
)

func GetArticlesWithKeyword(page uint64, pageSize uint64, keyword string) []models.Article {
	db := GetDB()
	articles := make([]models.Article, 0)
	fmt.Println(articles)
	db.
		Where("tags LIKE ?", "%"+keyword+"%").
		Or("title LIKE ?", "%"+keyword+"%").
		Offset(page * pageSize).
		Limit(pageSize).
		Find(&articles)
	fmt.Println(articles)

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
