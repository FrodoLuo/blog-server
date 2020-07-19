package main

import (
	"blog-server/resources/articles"
	"blog-server/resources/comments"
	"blog-server/resources/users"

	"github.com/gin-gonic/gin"
)

func main() {
	ginInstance := gin.Default()
	routerGroup := ginInstance.Group("/api")

	{
		routerGroup.GET("/articles", articles.GetArticleList)
		routerGroup.GET("/articles/detail/:id", articles.GetCertainArticle)
		routerGroup.GET("/articles/count", articles.CountArticle)
		routerGroup.POST("/articles", articles.Post)
	}

	{
		routerGroup.POST("/comments", comments.Post)
	}

	{
		routerGroup.GET("/users/:id", users.Get)
	}

	ginInstance.Run("127.0.0.1:3100")
}
