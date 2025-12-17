package main

import (
	"fmt"

	"go-blog/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	fmt.Println("App running on port:", cfg.AppPort)
	fmt.Println("DB host:", cfg.DBHost)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.Run(":" + cfg.AppPort)
}
