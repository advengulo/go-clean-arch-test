package domains

import (
	"time"
)

// UserID is the ID of a user
type UserID uint

// User contains information about a user
type User struct {
	ID        UserID     `gorm:"column:id" json:"-"`
	Username  string     `gorm:"column:username" json:"username"  validate:"required,min=3"`
	Password  string     `gorm:"column:password" json:"password" validate:"required,min=3"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"-"`
}
