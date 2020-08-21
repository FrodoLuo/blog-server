package router

import (
	"blog-server/middlewares"
	"blog-server/models"
	"blog-server/resources/articles"
	"blog-server/resources/comments"
	"blog-server/resources/configs"
	"blog-server/resources/media"
	"blog-server/resources/users"
	"os"

	"github.com/gin-gonic/gin"
)

func BindRouters(ginInstance *gin.Engine) {

	ginInstance.MaxMultipartMemory = 8 << 20

	cmsRoot := os.Getenv("BLOG_CMS_ROOT")
	if cmsRoot == "" {
		cmsRoot = "./public"
	}

	ginInstance.Static("/admin", cmsRoot)

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
		publicRoutes.GET("/comments", comments.Get)
		publicRoutes.GET("/comments/count", comments.Count)
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

	{ // media
		publicRoutes.GET("/media", media.Get)
		publicRoutes.GET("/media/count", media.Count)

		adminRequiredRoutes.POST("/media", media.Post)
	}
}
