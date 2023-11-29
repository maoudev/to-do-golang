package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maoudev/todo/internal/infraestructure/api/task"
	"github.com/maoudev/todo/internal/infraestructure/api/user"
)

func setupRoutes(e *gin.Engine) {
	user.SetupRoutes(e)
	task.SetupRoutes(e)
}
