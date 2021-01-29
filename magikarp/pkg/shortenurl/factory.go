package shortenurl

import (
	log "github.com/sirupsen/logrus"
	"github.com/system-design-url-shortener/magikarp/entity"
	"go.uber.org/dig"
)

// New is factory function for create shorten url service
func New(logger *log.Logger, urlBackend entity.URLBackend, cacheBackend entity.CacheBackend, config Config) (entity.ShortenURLService, error) {
	if err := urlBackend.InitSchema(); err != nil {
		return nil, err
	}
	return &service{
		Logger:       logger,
		Config:       config,
		URLBackend:   urlBackend,
		CacheBackend: cacheBackend,
	}, nil
}

// Config for shortenURL service.
type Config struct {
	dig.In
	URLLength int `name:"shortenurl_url_length"`
}
