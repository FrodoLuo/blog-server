package comments

import (
	"blog-server/models"
	"blog-server/models/responsemodel"
	"blog-server/services"

	"github.com/gin-gonic/gin"
)

// Post method of Comment
func Post(ctx *gin.Context) {
	commentToSave := models.Comment{}
	ctx.BindJSON(&commentToSave)
	if commentToSave.ArticleID == 0 ||
		commentToSave.Nickname == "" ||
		commentToSave.Content == "" {
		ctx.AbortWithStatusJSON(415, responsemodel.ErrorMessage{
			Message: "Empty params are not allowed",
		})
		return
	}
	services.CreateComment(&commentToSave)
	ctx.JSON(200, commentToSave)
}

func Get(ctx *gin.Context) {
	page, pageSize, err := services.ParsePageAndSize(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}
	comments := services.GetComments(uint(page), uint(pageSize))
	ctx.JSON(200, comments)
}

func Count(ctx *gin.Context) {
	ctx.JSON(200, services.CountComments())
}
