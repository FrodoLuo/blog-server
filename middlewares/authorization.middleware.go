package middlewares

import (
	"blog-server/models"
	"blog-server/services"

	"github.com/gin-gonic/gin"
)

/*
AuthorizationMiddleware create a middleware with configurable level to guard request
*/
func AuthorizationMiddleware(allowedRoles []models.UserRole) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
		if contains(allowedRoles, user.Type) {
			ctx.Next()
		} else {
			ctx.AbortWithStatus(403)
			return
		}
	}
}

func contains(validLevels []int, level int) bool {
	for _, item := range validLevels {
		if item == level {
			return true
		}
	}
	return false
}
