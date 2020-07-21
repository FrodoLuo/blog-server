package users

import (
	"blog-server/services"
	"fmt"
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

func GetWithToken(ctx *gin.Context) {
	token, err := ctx.Cookie("token")
	if err != nil {
		ctx.AbortWithStatusJSON(401, err)
		return
	}
	user := services.GetUserByToken(token)
	if user.ID == 0 {
		ctx.AbortWithStatus(401)
		return
	}
	ctx.JSON(200, user)
}

func Post(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	fmt.Println(email)
	fmt.Println(password)
	token, status := services.UpdateAuth(email, password)
	if status != 0 {
		ctx.AbortWithStatus(status)
		return
	}
	ctx.Header("Set-Cookie", "token="+token+"; httponly")
	ctx.Status(204)
}
