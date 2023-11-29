package mysql

import (
	"github.com/google/uuid"
	"github.com/maoudev/todo/internal/pkg/domain"
	"gorm.io/gorm"
)

// Client represents an instance to the database.
type client struct {
	db *gorm.DB
}

func NewClient() *client {
	return &client{
		db: connect(),
	}
}

func (c *client) Create(value interface{}) error {
	return c.db.Create(value).Error
}

func (c *client) First(dest interface{}, conds ...interface{}) error {
	return c.db.First(dest, conds...).Error
}

func (c *client) GetUserTasks(userID uuid.UUID) ([]*domain.Task, error) {
	tasks := []*domain.Task{}
	err := c.db.Find(&tasks, "user_id = ?", userID).Order("start_time asc").Error
	return tasks, err
}

func (c *client) ToggleStateTask(taskID uuid.UUID) error {
	task := &domain.Task{}
	if err := c.db.First(task, "id = ?", taskID).Error; err != nil {
		return err
	}

	var state bool

	if task.Active {
		state = false
	} else {
		state = true
	}

	return c.db.Model(&domain.Task{}).Where("id = ?", taskID).Update("active", state).Error
}

func (c *client) DeleteTask(taskID uuid.UUID) error {
	return c.db.Delete(&domain.Task{}, "id = ?", taskID).Error
}
