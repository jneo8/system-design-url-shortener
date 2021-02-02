package postgres

import (
	"github.com/google/uuid"
	"github.com/system-design-url-shortener/magikarp/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// MakePassword : Encrypt user password
func MakePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CompareHashAndPassword ....
func CompareHashAndPassword(hashPwd string, plainPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(plainPwd)); err != nil {
		return false
	}
	return true
}

// User ...
type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id;default:uuid_generate_v4()"`
	DisplayName string    `db:"display_name"`
	UserName    string    `db:"username" gorm:"unique"`
	Password    string    `db:"password"`
	Email       string    `db:"email"`
	IsStaff     bool      `db:"is_staff"`
	IsSuperUser bool      `db:"is_superuser"`
	IsActive    bool      `db:"is_active"`
	LastLogin   time.Time `db:"last_login"`
}

// BeforeCreate for User.
func (user *User) BeforeCreate(tx *gorm.DB) error {
	if user.Password != "" {
		hash, err := MakePassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hash
	}
	return nil
}

func (user *User) toEntityUser() entity.User {
	return entity.User{
		ID:          user.ID,
		DisplayName: user.DisplayName,
		UserName:    user.UserName,
		Password:    user.Password,
		Email:       user.Email,
		IsStaff:     user.IsStaff,
		IsSuperUser: user.IsSuperUser,
		IsActive:    user.IsActive,
		LastLogin:   user.LastLogin,
	}
}
