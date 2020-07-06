package articles

import (
	"blog-server/models"
	"blog-server/services"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	ctx.JSON(200, services.GetAllArticles())
}

func Post(ctx *gin.Context) {
	targetArticle := models.Article{}
	err := ctx.BindJSON(&targetArticle)
	services.CreateArticle(&targetArticle)
	if err != nil {
		ctx.JSON(415, err)
	}
	ctx.JSON(200, targetArticle)
}
