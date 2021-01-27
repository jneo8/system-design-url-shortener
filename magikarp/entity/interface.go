package entity

import (
	"github.com/google/uuid"
)

// ShortenURLService is service handle shorten url.
type ShortenURLService interface {
	// Return new shorten url.
	ShortenURL(
		originalURL string,
		expireTime int64,
		userID uuid.UUID,
	) (url URL, err error)

	// Delete url from DB.
	DeleteURL(
		urlKey uuid.UUID,
	) (err error)
}

// URLBackend is repository interface for all db backend.
type URLBackend interface {
	NewURL(url URL) (URL, error)
	DeleteURL(key uuid.UUID) error
	GetURLByShortURL(shortURL string) (url URL, err error)
	InitSchema() error
}
