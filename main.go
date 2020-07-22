package main

import (
	"blog-server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	ginInstance := gin.Default()
	router.BindRouters(ginInstance)
	ginInstance.Run("127.0.0.1:3100")
}
