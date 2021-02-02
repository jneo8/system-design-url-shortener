package postgres

import (
	"github.com/google/uuid"
	"github.com/system-design-url-shortener/magikarp/entity"
	"gorm.io/gorm"
)

// URL ...
type URL struct {
	gorm.Model
	UserID      *uuid.UUID
	OriginalURL string
	ShortURL    string
	ExpireTime  int64
}

func (url *URL) toEntityURL() entity.URL {
	return entity.URL{
		UserID:      url.UserID,
		URLID:       url.Model.ID,
		OriginalURL: url.OriginalURL,
		ShortURL:    url.ShortURL,
		ExpireTime:  url.ExpireTime,
		CreateTime:  url.Model.CreatedAt.UnixNano()}
}
