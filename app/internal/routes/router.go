package routes

import (
	"app/configs"
	"app/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(secretKey string, router *gin.Engine, handlers *configs.Handlers) {
	publicRoutes := router.Group("")
	{
		// Authentication Routes
		publicRoutes.POST("/register", handlers.AuthHandler.Register)
		publicRoutes.POST("/login", handlers.AuthHandler.Login)

		// Post Routes
		publicRoutes.GET("/posts/:id", handlers.PostHandler.GetPostByID)
		publicRoutes.GET("/posts", handlers.PostHandler.GetAllPosts)
	}

	privateRoutes := router.Group("")
	privateRoutes.Use(middlewares.AuthMiddleware(secretKey))
	{
		// Posts Routes
		privateRoutes.POST("/posts", handlers.PostHandler.CreatePost)

		// Comments Routes
	}

}
