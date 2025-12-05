package main

import (
	"AdvAuthGo/config"
	"AdvAuthGo/internal/database"
	"AdvAuthGo/internal/handlers"
	"AdvAuthGo/internal/repositories"
	"AdvAuthGo/internal/services"
	"AdvAuthGo/internal/utils"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)
	emailSender := utils.NewEmailSender(cfg)

	database.Migrate(db)

	userRepo := repositories.NewUserRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)

	authService := services.NewAuthService(userRepo, tokenRepo, emailSender, cfg)
	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
		api.GET("/activate/:token", authHandler.Activate)
		api.POST("/refresh", authHandler.Refresh)
		api.GET("/users", authHandler.GetUsers)
	}

	log.Printf("Server started on port %s", cfg.ServerPort)
	router.Run(":8080")
}
