package domain

import (
	"time"

	"github.com/google/uuid"
)

// Task represents a user task.
type Task struct {
	ID uuid.UUID `gorm:"unique"`

	Name        string `json:"name" gorm:"varchar:100" binding:"required"`
	Description string `json:"description" gorm:"type:text"`

	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	UserID uuid.UUID `json:"userID" gorm:"size:300"`

	Active bool `gorm:"default:true"`

	CreatedAt time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true"`
}
