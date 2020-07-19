package responsemodel

import (
	"blog-server/models"
	"time"
)

type ArticleResponse struct {
	Title     string      `json:"title"`
	UpdatedAt time.Time   `json:"updatedAt"`
	Author    models.User `json:"author"`
	Brief     string      `json:"brief"`
	Tags      string      `json:"tags"`
}
