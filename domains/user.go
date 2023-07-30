package domains

import (
	"time"
)

// User contains information about a user
type User struct {
	ID        uint       `gorm:"column:id" json:"id"`
	Username  string     `gorm:"column:username" json:"username"  validate:"required,min=3"`
	Password  string     `gorm:"column:password" json:"password" validate:"required,min=3"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

type ValidateUserUpdate struct {
	ID          uint   `json:"id" validate:"required"`
	Username    string `json:"username" validate:"required,min=3"`
	NewPassword string `json:"new_password" validate:"required,min=3"`
	OldPassword string `json:"old_password" validate:"required,min=3"`
}
