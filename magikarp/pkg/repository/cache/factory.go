package cache

import (
	"github.com/patrickmn/go-cache"
	"github.com/system-design-url-shortener/magikarp/entity"
	"time"
)

// New is a simple factory func return CacheBackend.
func New() (entity.CacheBackend, error) {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &repo{Cache: c}, nil
}
