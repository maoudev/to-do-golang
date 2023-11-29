package domain

import "github.com/google/uuid"

// User represents a normal application user.
type User struct {
	ID       uuid.UUID `gorm:"unique; size:300"` // ID is the unique identifier of each user.
	Name     string    `json:"name" gorm:"varchar:50" binding:"required"`
	LastName string    `json:"last_name" gorm:"varchar:50" binding:"required"`
	Email    string    `json:"email" binding:"required" gorm:"unique"`
	Password string    `json:"password" binding:"required"`

	Tasks []Task `gorm:"foreignKey:UserID"`
}
