package models

import (
	"time"
)

type Media struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Tag            string    `json:"tag"`
	Description    string    `json:"description"`
	Source         string    `json:"source"`
	OrderReference uint      `json:"orderReference"`
}
