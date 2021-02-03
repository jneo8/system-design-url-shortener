package api

import (
	"github.com/system-design-url-shortener/magikarp/entity"
)

// SignupForm for user signup.
type SignupForm struct {
	UserName       string `form:"userName" binding:"required"`
	DisplayName    string `form:"displayName"`
	Password       string `form:"password" binding:"required"`
	SecondPassword string `form:"secondPassword" binding:"required"`
	Email          string `form:"email" binding:"required"`
}

type shortenerFuncForm struct {
	APIDevKey string `form:"apiDevKey"`
	// UserID      string `form:"userID"`
	OriginalURL string `form:"originalURL" binding:"required"`
	ExpireTime  int64  `form:"expireTime"`
}

// ToUser transform form to user.
func (form *SignupForm) ToUser() entity.User {
	return entity.User{
		UserName:    form.UserName,
		Password:    form.Password,
		DisplayName: form.Password,
		Email:       form.Email,
	}
}

// ValidPassword ...
func (form *SignupForm) ValidPassword() bool {
	return form.Password == form.SecondPassword
}
