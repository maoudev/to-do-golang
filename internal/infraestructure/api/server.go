package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/maoudev/todo/internal/config"
)

func RunServer() {
	server := gin.Default()

	server.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		MaxAge:         50 * time.Second,
	}))

	server.GET("/api/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Bienvenido a la api",
		})
	})

	gin.SetMode(gin.ReleaseMode)

	setupRoutes(server)

	port := fmt.Sprintf(":%v", config.API_PORT)
	server.Run(port)
}
