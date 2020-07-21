package middlewares

import (
	"blog-server/models"

	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(allowedRoles []models.UserRole) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
