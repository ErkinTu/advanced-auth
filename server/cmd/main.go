package main

import (
	"AdvAuthGo/config"
	"AdvAuthGo/internal/database"
	"AdvAuthGo/internal/handlers"
	"AdvAuthGo/internal/repositories"
	"AdvAuthGo/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)

	database.Migrate(db)

	userRepo := repositories.NewUserRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)

	authService := services.NewAuthService(userRepo, tokenRepo, cfg)
	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
		api.POST("/activate/:token", authHandler.Activate)
		api.POST("/refresh/:token", authHandler.Refresh)
		api.GET("/users", authHandler.GetUsers)
	}

	log.Println("Server started on port %s", cfg.ServerPort)
	router.Run(":8080")
}
