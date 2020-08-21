package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Content   string    `json:"content"`
	Permitted bool      `json:"permitted"`
	ArticleID uint      `json:"articleId"`
	Article   Article   `gorm:"foreignKey:ArticleID" json:"article"`
}
