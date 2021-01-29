package postgres

import (
	"github.com/google/uuid"
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
