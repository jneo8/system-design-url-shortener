package postgres

import (
	"github.com/system-design-url-shortener/magikarp/entity"
)

func (r *repo) NewUser(entityUser entity.User) (entity.User, error) {
	user := entityUserToUser(entityUser)
	if result := r.DB.Create(&user); result.Error != nil {
		return entity.User{}, result.Error
	}
	return user.toEntityUser(), nil
}

func (r *repo) UserLogin(entityUser entity.User) (entity.User, bool) {
	dbUser := User{}
	if result := r.DB.Where(&User{UserName: entityUser.UserName}).Take(&dbUser); result.Error != nil {
		r.Logger.Error(result.Error)
		return entityUser, false
	}
	if !CompareHashAndPassword(dbUser.Password, entityUser.Password) {
		return entityUser, false
	}
	return dbUser.toEntityUser(), true
}
