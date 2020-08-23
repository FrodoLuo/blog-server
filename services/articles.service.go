package services

import (
	"blog-server/models"
)

func GetArticlesWithKeyword(page uint64, pageSize uint64, keyword string) []models.Article {
	db := GetDB()
	articles := make([]models.Article, 0)
	db.
		Preload("Author").
		Preload("Comments").
		Where("Tags LIKE ?", "%"+keyword+"%").
		Or("Title LIKE ?", "%"+keyword+"%").
		Order("updated_at DESC", true).
		Offset(page * pageSize).
		Limit(pageSize).
		Find(&articles)

	return articles
}

func GetArticleWithID(id uint64) *models.Article {
	db := GetDB()
	article := models.GenerateEmptyArticle()
	db.
		Preload("Author").
		Preload("Comments").
		Where("id = ?", id).
		Find(&article)

	return article
}

func CountArticlesWithKeyword(keyword string) uint {
	db := GetDB()
	var count uint
	db.
		Model(&models.Article{}).
		Where("Tags LIKE ?", "%"+keyword+"%").
		Or("Title LIKE ?", "%"+keyword+"%").
		Count(&count)
	return count
}

/*
UpdateOrCreateArticle method to create and return the entity of the a article
*/
func UpdateOrCreateArticle(articleToSave models.Article, userID uint) models.Article {
	if articleToSave.AuthorID == 0 {
		articleToSave.AuthorID = userID
	}
	GetDB().
		Preload("Author").
		Save(&articleToSave).
		Find(&articleToSave)
	return articleToSave
}
