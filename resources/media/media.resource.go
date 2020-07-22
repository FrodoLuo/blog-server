package media

import (
	"blog-server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	tag := ctx.Query("tag")
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "0"))
	if err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}
	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}
	ctx.JSON(200, services.GetMediaByTag(tag, uint(page), uint(pageSize)))
}

func Post(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	tag := ctx.PostForm("tag")
	description := ctx.PostForm("description")
	order, err := strconv.Atoi(ctx.DefaultPostForm("order", "0"))

	if err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	savedMedia := services.SaveMedia(file, tag, description, uint(order))
	ctx.JSON(200, savedMedia)
}
