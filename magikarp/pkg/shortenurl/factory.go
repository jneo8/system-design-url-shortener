package shortenurl

import (
	"github.com/google/uuid"
	"github.com/system-design-url-shortener/magikarp/entity"
)

// New is factory function for create shorten url service
func New() (entity.ShortenURLService, error) {
	return &service{}, nil
}

type service struct {
}

func (s *service) ShortenURL(originalURL string, expireTime int64, userID uuid.UUID) (entity.URL, error) {
	return entity.URL{ShortURL: "abc.com"}, nil
}

func (s *service) DeleteURL(urlKey uuid.UUID) error {
	return nil
}
