package entity

import (
	"github.com/google/uuid"
)

// URL ...
type URL struct {
	UserUUID    uuid.UUID // Creater.
	Key         uuid.UUID // Primary key.
	OriginalURL string
	ShortenURL  string
	ExpireDate  int64
	CreateTime  int64
}
