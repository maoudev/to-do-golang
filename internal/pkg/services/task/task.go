package task

import (
	"errors"
	"time"

	"github.com/maoudev/todo/internal/pkg/domain"
	"github.com/maoudev/todo/internal/pkg/ports"
	"github.com/maoudev/todo/internal/pkg/utils"
)

var (
	ErrInvalidDateRange = errors.New("the date range is invalid")
	ErrInvalidDate      = errors.New("the dates entered are invalid")
	ErrInvalidUserID    = errors.New("the user ID granted is invalid")
)

type taskService struct {
	repository ports.TaskRepository
}

func NewService(repository ports.TaskRepository) *taskService {
	return &taskService{
		repository: repository,
	}
}

func (t *taskService) Create(task *domain.Task, startDate, endDate string, userID string) error {
	var err error

	task.ID = utils.CreateID()
	task.Active = true

	if !isDateRangeValid(startDate, endDate) {
		return ErrInvalidDateRange
	}

	task.StartTime, task.EndTime, err = parseDateRange(startDate, endDate)
	if err != nil {
		return ErrInvalidDate
	}

	task.UserID, err = utils.ParseUUID(userID)
	if err != nil {
		return ErrInvalidUserID
	}

	return t.repository.Create(task)
}

func (t *taskService) GetTasks(userID string) ([]*domain.Task, error) {
	parsedID, err := utils.ParseUUID(userID)
	if err != nil {
		return nil, ErrInvalidUserID
	}

	tasks, err := t.repository.GetUserTasks(parsedID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *taskService) GetTaskById(taskID string) (*domain.Task, error) {
	parsedID, err := utils.ParseUUID(taskID)
	if err != nil {
		return nil, err
	}

	task := &domain.Task{}

	if err := t.repository.First(task, "id = ?", parsedID); err != nil {
		return nil, err
	}

	return task, nil
}

func (t *taskService) DeleteTaskByID(taskID string) error {
	parsedID, err := utils.ParseUUID(taskID)
	if err != nil {
		return err
	}

	if err := t.repository.DeleteTask(parsedID); err != nil {
		return err
	}

	return nil
}

func (t *taskService) MarkTask(taskID string) error {
	parsedID, err := utils.ParseUUID(taskID)
	if err != nil {
		return err
	}

	if err := t.repository.ToggleStateTask(parsedID); err != nil {
		return err
	}

	return nil
}

func parseDateRange(startDate, endDate string) (time.Time, time.Time, error) {
	st, err := time.Parse("2006-1-2", startDate)
	if err != nil {
		return time.Now(), time.Now(), err
	}

	ed, err := time.Parse("2006-1-2", endDate)
	if err != nil {
		return time.Now(), time.Now(), err
	}

	return st, ed, nil
}

// IsDateRangeValid checks if the date range is valid.
func isDateRangeValid(startDate, endDate string) bool {
	if startDate == "" || endDate == "" {
		return false
	}

	st, ed, err := parseDateRange(startDate, endDate)
	if err != nil {
		return false
	}

	return ed.After(st)
}
