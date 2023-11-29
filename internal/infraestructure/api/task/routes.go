package task

import (
	"github.com/gin-gonic/gin"
	"github.com/maoudev/todo/internal/infraestructure/api/middlewares"
	"github.com/maoudev/todo/internal/infraestructure/repositories/mysql"
	"github.com/maoudev/todo/internal/pkg/services/task"
)

func SetupRoutes(e *gin.Engine) {
	repo := mysql.NewClient()
	service := task.NewService(repo)
	handler := newHandler(service)

	v1 := e.Group("/api/v1")
	{
		v1.POST("/task", middlewares.Authenticate(), handler.CreateTask)
		v1.GET("/task", middlewares.Authenticate(), handler.Get)
		v1.GET("/task/:id", middlewares.Authenticate(), handler.GetTaskById)
		v1.PUT("/task/:id", middlewares.Authenticate(), handler.MarkTask)
		v1.DELETE("/task/:id", middlewares.Authenticate(), handler.Delete)

	}

}
