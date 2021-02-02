package entity

import (
	"github.com/google/uuid"
	"time"
)

// User ...
type User struct {
	ID          uuid.UUID `db:"id" gorm:"primaryKey"`
	DisplayName string    `db:"display_name"`
	UserName    string    `db:"username" gorm:"unique"`
	Password    string    `db:"password"`
	Email       string    `db:"email"`
	IsStaff     bool      `db:"is_staff"`
	IsSuperUser bool      `db:"is_superuser"`
	IsActive    bool      `db:"is_active"`
	LastLogin   time.Time `db:"last_login"`
}
