package cache

import (
	"github.com/patrickmn/go-cache"
	"github.com/system-design-url-shortener/magikarp/entity"
)

type repo struct {
	Cache *cache.Cache
}

func (r *repo) Get(shortURL string) (string, bool) {
	if url, getCacheOK := r.Cache.Get(shortURL); getCacheOK {
		if urlStr, ok := url.(string); ok {
			return urlStr, true
		}
	}
	return "", false
}

func (r *repo) Set(url entity.URL) error {
	r.Cache.Set(url.ShortURL, url.OriginalURL, cache.DefaultExpiration)
	return nil
}
