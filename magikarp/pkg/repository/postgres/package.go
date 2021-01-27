package postgres

import (
	"github.com/google/uuid"
	"github.com/system-design-url-shortener/magikarp/entity"
)

// New is a simple factory func return URLBackend.
func New() entity.URLBackend {
	return &repo{}
}

type repo struct{}

func (r *repo) InitSchema() error {
	return nil
}

func (r *repo) NewURL(url entity.URL) (entity.URL, error) {
	return entity.URL{}, nil
}

func (r *repo) DeleteURL(key uuid.UUID) error {
	return nil
}

func (r *repo) GetURLByShortURL(shortURL string) (entity.URL, error) {
	return entity.URL{}, nil
}
