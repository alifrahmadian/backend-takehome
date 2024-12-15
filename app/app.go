package main

import (
	"app/configs"
	"app/internal/db"
	"app/internal/handlers"
	"app/internal/repositories"
	"app/internal/routes"
	"app/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Config *configs.Config
}

func LoadConfig() (*configs.Config, error) {
	err := configs.LoadGoDotEnv()
	if err != nil {
		return nil, err
	}

	dbConfig := configs.LoadDBConfig()
	env := configs.LoadEnv()
	authConfig := configs.LoadAuthConfig()

	db, err := db.Connect(*dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)

	authService := services.NewAuthService(userRepo)
	postService := services.NewPostService(postRepo)

	authHandler := handlers.NewAuthHandler(&authService, authConfig.SecretKey, authConfig.TTL)
	postHandler := handlers.NewPostHandler(&postService)

	return &configs.Config{
		DB:   db,
		Env:  env,
		Auth: authConfig,
		Handlers: &configs.Handlers{
			AuthHandler: authHandler,
			PostHandler: postHandler,
		},
	}, nil
}

func NewApp() *App {
	cfg, err := LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	router := gin.Default()
	routes.SetupRoutes(cfg.Auth.SecretKey, router, cfg.Handlers)

	return &App{
		Router: router,
		Config: cfg,
	}
}
