package routes

import (
	"app/configs"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handlers *configs.Handlers) {
	router.POST("/register", handlers.AuthHandler.Register)
	router.POST("/login", handlers.AuthHandler.Login)
}
