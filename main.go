package main

import (
	"blog-server/resources/articles"

	"github.com/gin-gonic/gin"
)

func main() {
	ginInstance := gin.Default()
	routerGroup := ginInstance.Group("/api")

	routerGroup.GET("/articles", articles.Get)

	ginInstance.Run("127.0.0.1:3100")
}
