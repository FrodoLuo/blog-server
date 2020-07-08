package main

import (
	"blog-server/resources/articles"
	"blog-server/resources/users"

	"github.com/gin-gonic/gin"
)

func main() {
	ginInstance := gin.Default()
	routerGroup := ginInstance.Group("/api")

	{
		routerGroup.GET("/articles", articles.GetArticleList)
		routerGroup.GET("/articles/:id", articles.GetCertainArticle)
		routerGroup.POST("/articles", articles.Post)
	}

	{
		routerGroup.GET("/users/:id", users.Get)
	}

	ginInstance.Run("127.0.0.1:3100")
}
