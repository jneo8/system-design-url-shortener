package postgres

import (
	log "github.com/sirupsen/logrus"
	"github.com/system-design-url-shortener/magikarp/entity"
	"gorm.io/gorm"
)

type repo struct {
	Config Config
	DB     *gorm.DB
	Logger *log.Logger
}

func (r *repo) InitSchema() error {
	if err := r.DB.AutoMigrate(&URL{}); err != nil {
		return err
	}
	return nil
}

func (r *repo) NewURL(url entity.URL) (entity.URL, error) {
	newURL := entityURLToURL(url)
	if url.EmptyUser() {
		if result := r.DB.Create(&newURL); result.Error != nil {
			return entity.URL{}, result.Error
		}
	} else {
		if result := r.DB.FirstOrCreate(&newURL, URL{UserID: newURL.UserID, OriginalURL: newURL.OriginalURL}); result.Error != nil {
			return entity.URL{}, result.Error
		}
	}

	return newURL.toEntityURL(), nil
}

func (r *repo) DeleteURL(urlID int64) error {
	return nil
}

func (r *repo) GetURLByShortURL(shortURL string) (entity.URL, error) {
	var url URL
	if err := r.DB.Where(&URL{ShortURL: shortURL}).First(&url).Error; err != nil {
		return entity.URL{}, err
	}
	return url.toEntityURL(), nil
}
