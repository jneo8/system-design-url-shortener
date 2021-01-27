package entity

import (
	"github.com/google/uuid"
)

// Service Interface.
type Service interface {
	// Return new shorten url.
	ShortenURL(
		token string,
		originalURL string,
		expireDate int64,
		userUUID uuid.UUID,
		customAlias map[string]string,
	) (url URL, err error)

	// Delete url from DB.
	DeleteURL(
		token string,
		urlKey uuid.UUID,
	) (err error)
}
