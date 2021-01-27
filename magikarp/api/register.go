package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/system-design-url-shortener/magikarp/entity"
)

// Single tone logger for api.
var logger *log.Logger
var apiDevKey string

// RegisterAPI for API endpoints.
func RegisterAPI(newLogger *log.Logger, newAPIDevKey string, service entity.ShortenURLService) (*gin.Engine, error) {

	logger = newLogger
	apiDevKey = newAPIDevKey
	logger.Infof("RegisterAPI")

	r := gin.New()

	// Ping/Pong.
	r.GET("/ping", pingFunc())

	// shorten url
	r.POST("/url", shortenerFunc(service))

	return r, nil
}
