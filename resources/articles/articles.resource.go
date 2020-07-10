package articles

import (
	"blog-server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticleList(ctx *gin.Context) {
	page, err := strconv.ParseUint(ctx.DefaultQuery("page", "0"), 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}
	pageSize, err := strconv.ParseUint(ctx.DefaultQuery("pageSize", "10"), 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}
	keyword := ctx.DefaultQuery("keyword", "")
	ctx.JSON(200, services.GetArticlesWithKeyword(page, pageSize, keyword))
}

func Post(ctx *gin.Context) {
	targetArticle := services.NewArticleParams{}
	err := ctx.BindJSON(&targetArticle)
	savedArticle := services.CreateArticle(&targetArticle)
	if err != nil {
		ctx.JSON(415, err)
	}
	ctx.JSON(200, savedArticle)
}

func GetCertainArticle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}
	ctx.JSON(200, services.GetArticleWithID(id))
}
