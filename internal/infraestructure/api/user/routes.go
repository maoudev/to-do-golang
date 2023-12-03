package user

import (
	"github.com/gin-gonic/gin"
	"github.com/maoudev/todo/internal/infraestructure/api/middlewares"
	"github.com/maoudev/todo/internal/infraestructure/repositories/mysql"
	"github.com/maoudev/todo/internal/pkg/services/user"
)

func SetupRoutes(e *gin.Engine) {
	repo := mysql.NewClient()
	service := user.NewService(repo)
	handler := newHandler(service)

	v1 := e.Group("/api/v1")
	{
		v1.POST("/user", handler.CreateUser)
		v1.POST("/authenticate", handler.Login)
		v1.GET("/user", middlewares.Authenticate(), handler.GetUser)
	}

}
