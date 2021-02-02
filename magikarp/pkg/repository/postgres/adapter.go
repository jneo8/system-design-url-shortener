package postgres

import (
	"github.com/system-design-url-shortener/magikarp/entity"
)

func entityURLToURL(url entity.URL) URL {
	return URL{
		UserID:      url.UserID,
		OriginalURL: url.OriginalURL,
		ShortURL:    url.ShortURL,
		ExpireTime:  url.ExpireTime,
	}
}

func entityUserToUser(user entity.User) User {
	return User{
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
