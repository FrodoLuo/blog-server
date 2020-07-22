package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ginInstance := gin.Default()
	BindRouters(ginInstance)
	ginInstance.Run("127.0.0.1:3100")
}
