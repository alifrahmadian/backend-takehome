package routes

import (
	"app/configs"
	"app/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(secretKey string, router *gin.Engine, handlers *configs.Handlers) {
	postsRoutes := router.Group("/posts")
	// commentRoutes := router.Group("/posts/:id/comments")

	publicRoutes := router.Group("")
	{
		// Authentication Routes
		publicRoutes.POST("/register", handlers.AuthHandler.Register)
		publicRoutes.POST("/login", handlers.AuthHandler.Login)

		// Post Routes
		postsRoutes.GET("/:id", handlers.PostHandler.GetPostByID)
	}

	privateRoutes := router.Group("")
	privateRoutes.Use(middlewares.AuthMiddleware(secretKey))
	{
		// Posts Routes
		postsRoutes.POST("", handlers.PostHandler.CreatePost)

		// Comments Routes
	}

}
