package services

import (
	"blog-server/models"
)

func GetArticlesWithKeyword(page uint64, pageSize uint64, keyword string) []models.Article {
	db := GetDB()
	articles := make([]models.Article, 0)
	db.
		Where("tags LIKE ?", "%"+keyword+"%").
		Or("title LIKE ?", "%"+keyword+"%").
		Offset(page * pageSize).
		Limit(pageSize).
		Find(&articles)

	return articles
}

func GetArticleWithID(id uint64) *models.Article {
	db := GetDB()
	article := models.GenerateEmptyArticle()
	db.
		Where("id = ?", id).
		Find(&article)

	return article
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
	articleToSave.AuthorID = articleParams.AuthorID

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
