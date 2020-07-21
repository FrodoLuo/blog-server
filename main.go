package main

import (
	"blog-server/middlewares"
	"blog-server/models"
	"blog-server/resources/articles"
	"blog-server/resources/comments"
	"blog-server/resources/configs"
	"blog-server/resources/users"

	"github.com/gin-gonic/gin"
)

func main() {
	ginInstance := gin.Default()
	publicRoutes := ginInstance.Group("/api")
	authRequiredRoutes := ginInstance.
		Group("/api").
		Use(middlewares.AuthorizationMiddleware(
			[]models.UserRole{
				models.REGISTERED,
				models.ADMIN,
			},
		))

	adminRequiredRoutes := ginInstance.
		Group("/api").
		Use(middlewares.AuthorizationMiddleware(
			[]models.UserRole{
				models.ADMIN,
			},
		))

	{ // articles
		publicRoutes.GET("/articles", articles.GetArticleList)
		publicRoutes.GET("/articles/detail/:id", articles.GetCertainArticle)
		publicRoutes.GET("/articles/count", articles.CountArticle)

		authRequiredRoutes.POST("/articles", articles.Post)
	}

	{ // comments
		publicRoutes.POST("/comments", comments.Post)
	}

	{ // configs
		publicRoutes.GET("/configs", configs.GetAll)
		publicRoutes.GET("/configs/detail/:title", configs.Get)

		authRequiredRoutes.POST("/configs", configs.Post)
	}

	{ // users
		publicRoutes.POST("/users", users.Post)
		publicRoutes.GET("/users", users.GetWithToken)

		adminRequiredRoutes.GET("/users/:id", users.Get)
	}

	ginInstance.Run("127.0.0.1:3100")
}
