package shortenurl

import (
	log "github.com/sirupsen/logrus"
	"github.com/system-design-url-shortener/magikarp/entity"
	"go.uber.org/dig"
)

// New is factory function for create shorten url service
func New(logger *log.Logger, backend entity.Backend, cacheBackend entity.CacheBackend, config Config) (entity.ShortenURLService, error) {
	if err := backend.InitSchema(); err != nil {
		return nil, err
	}
	return &service{
		Logger:       logger,
		Config:       config,
		Backend:      backend,
		CacheBackend: cacheBackend,
	}, nil
}

// Config for shortenURL service.
type Config struct {
	dig.In
	URLLength int `name:"shortenurl_url_length"`
}
