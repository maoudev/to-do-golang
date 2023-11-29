package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func RunServer() {
	server := gin.Default()

	server.Use(cors.Middleware(cors.Config{
		Origins:        "to-do-golang-nspl-git-main-maou-boop.vercel.app",
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

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	server.Run(port)
}
