package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maoudev/todo/internal/pkg/domain"
	"github.com/maoudev/todo/internal/pkg/ports"
)

type taskHandler struct {
	service ports.TaskService
}

func newHandler(service ports.TaskService) *taskHandler {
	return &taskHandler{
		service: service,
	}
}

func (t *taskHandler) CreateTask(c *gin.Context) {
	startDate := c.Query("sdate")
	endDate := c.Query("edate")

	userID := c.MustGet("userID").(string)

	task := &domain.Task{}

	if err := c.BindJSON(task); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := t.service.Create(task, startDate, endDate, userID); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (t *taskHandler) Get(c *gin.Context) {
	userID := c.MustGet("userID").(string)

	tasks, err := t.service.GetTasks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, tasks)
}

func (t *taskHandler) GetTaskById(c *gin.Context) {
	taskID := c.Param("id")

	task, err := t.service.GetTaskById(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskHandler) MarkTask(c *gin.Context) {
	taskID := c.Param("id")

	if err := t.service.MarkTask(taskID); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (t *taskHandler) Delete(c *gin.Context) {
	taskID := c.Param("id")

	if err := t.service.DeleteTaskByID(taskID); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
