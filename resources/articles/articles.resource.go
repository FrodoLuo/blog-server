package articles

import (
	"blog-server/services"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	ctx.JSON(200, services.GetAllArticles())
}
