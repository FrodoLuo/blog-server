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
func CreateArticle(articleParams *NewArticleParams) *models.Article {
	db := GetDB()
	articleToSave := models.GenerateEmptyArticle()

	articleToSave.Content = articleParams.Content
	articleToSave.Title = articleParams.Title
	articleToSave.Tags = articleParams.Tags
	articleToSave.Brief = articleParams.Brief

	if db.NewRecord(&articleToSave) {
		// articleToSave.AuthorID = 1
		db.Create(&articleToSave)
	}
	return articleToSave
}

type NewArticleParams struct {
	Content  string `default:""`
	Title    string `default:""`
	Brief    string `default:""`
	Tags     string `default:"Tech"`
	AuthorID uint
}
