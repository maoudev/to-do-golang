package ports

import (
	"github.com/google/uuid"
	"github.com/maoudev/todo/internal/pkg/domain"
)

type TaskRepository interface {
	Create(value interface{}) error
	First(out interface{}, conds ...interface{}) error
	GetUserTasks(userID uuid.UUID) ([]*domain.Task, error)
	ToggleStateTask(taskID uuid.UUID) error
	DeleteTask(taskID uuid.UUID) error
}

type TaskService interface {
	Create(task *domain.Task, startDate, endDate string, userID string) error
	GetTasks(userID string) ([]*domain.Task, error)
	GetTaskById(taskID string) (*domain.Task, error)
	MarkTask(taskID string) error
	DeleteTaskByID(taskID string) error
}
