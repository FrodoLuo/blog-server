package articles

import (
	"blog-server/models"
	"blog-server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetArticleList Get method, GetArticleList return the list of articles with given page and pagesize
func GetArticleList(ctx *gin.Context) {
	page, pageSize, err := services.ParsePageAndSize(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}
	keyword := ctx.DefaultQuery("keyword", "")
	ctx.JSON(200, services.GetArticlesWithKeyword(page, pageSize, keyword))
}

// Post method, create or update articles
func Post(ctx *gin.Context) {
	articleToSave := models.Article{}
	err := ctx.BindJSON(&articleToSave)
	if err != nil {
		ctx.JSON(415, err)
	}
	userID := ctx.Value("userID").(uint)
	if articleToSave.ID != 0 && userID != articleToSave.AuthorID {
		ctx.AbortWithStatus(403)
		return
	}
	savedArticle := services.UpdateOrCreateArticle(articleToSave, uint(userID))
	ctx.JSON(200, savedArticle)
}

// GetCertainArticle Get method, return an article with id or throw 404
func GetCertainArticle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}
	ctx.JSON(200, services.GetArticleWithID(id))
}

// CountArticle Get method, return count of articles with given keyword
func CountArticle(ctx *gin.Context) {
	keyword := ctx.DefaultQuery("keyword", "")
	ctx.JSON(200, services.CountArticlesWithKeyword(keyword))
}
