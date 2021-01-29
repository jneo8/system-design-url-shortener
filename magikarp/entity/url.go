package entity

import (
	"github.com/google/uuid"
)

// URL ...
type URL struct {
	UserID      *uuid.UUID
	URLID       uint
	OriginalURL string
	ShortURL    string
	ExpireTime  int64
	CreateTime  int64
}
