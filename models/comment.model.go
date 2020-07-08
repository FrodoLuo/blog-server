package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Nickname  string    `json:"nickname"`
	Content   string    `json:"content"`
	Permitted bool      `json:"permitted"`
}