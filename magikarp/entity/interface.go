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
		userID *uuid.UUID,
	) (url URL, err error)

	GetByShortURL(
		url string,
	) (originalURL string, err error)

	// Delete url from DB.
	DeleteURL(
		urlID int64,
	) (err error)
}

// URLBackend is repository interface for all db backend.
type URLBackend interface {
	NewURL(url URL) (URL, error)
	DeleteURL(urlID int64) error
	GetURLByShortURL(shortURL string) (url URL, err error)
	InitSchema() error
}

// CacheBackend is interface for shortenURL cache.
type CacheBackend interface {
	Set(url URL) error
	Get(url string) (string, bool)
}
