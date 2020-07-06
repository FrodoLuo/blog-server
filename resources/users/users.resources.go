package users

import (
	"blog-server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, err)
	}
	user := services.GetUserById(uint(userID))
	if &user == nil {
		ctx.Status(404)
	}
	ctx.JSON(200, user)
}
