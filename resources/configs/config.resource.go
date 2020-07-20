package configs

import (
	"blog-server/models"
	"blog-server/services"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	title := ctx.Params.ByName("title")
	config := services.GetConfigByTitle(title)

	if config.ID == 0 {
		ctx.JSON(404, nil)
	} else {
		ctx.JSON(200, config)
	}
}

func GetAll(ctx *gin.Context) {
	configs := services.GetAllConfig()

	ctx.JSON(200, configs)
}

func Post(ctx *gin.Context) {
	config := models.Config{}

	if err := ctx.BindJSON(&config); err != nil {
		ctx.AbortWithStatusJSON(400, err)
		return
	}

	configSaved := services.CreateConfig(&config)
	ctx.JSON(200, configSaved)

}
