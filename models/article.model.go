package models

import (
	"time"
)

type Article struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Brief     string     `json:"brief"`
	Class     string     `json:"class"`
	Content   string     `json:"content"`
	Tags      string     `json:"tags"`
	Title     string     `json:"title"`
	Author    *User      `gorm:"foreignKey:AuthorID" gorm:"default:nil" json:"author"`
	Comments  []*Comment `json:"comments"`
}
