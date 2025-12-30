package main

import (
	"fmt"
	"log"

	"go-blog/internal/config"
	"go-blog/internal/database"
	"go-blog/internal/handler"
	"go-blog/internal/middleware"
	"go-blog/internal/seed"
	"go-blog/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	fmt.Println("App running on port:", cfg.AppPort)
	fmt.Println("DB host:", cfg.DBHost)

	// Initialize JWT secret
	utils.SetJWTSecret(cfg.JWTSecret)

	// Initialize database
	db, err := database.InitPostgres(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Seed admin user (for development/demo only)
	if err := seed.SeedAdmin(db, cfg); err != nil {
		log.Printf("Warning: Failed to seed admin user: %v", err)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Register auth routes
	authHandler := handler.NewAuthHandler(db)
	authGroup := r.Group("/auth")
	authHandler.RegisterRoutes(authGroup)

	// Public routes
	postHandler := handler.NewPostHandler(db)
	publicGroup := r.Group("/api")
	postHandler.RegisterPublicRoutes(publicGroup)

	// Protected routes
	userHandler := handler.NewUserHandler(db)
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.AuthMiddleware())
	apiGroup.GET("/me", userHandler.GetMe)
	postHandler.RegisterRoutes(apiGroup)

	r.Run(":" + cfg.AppPort)
}
