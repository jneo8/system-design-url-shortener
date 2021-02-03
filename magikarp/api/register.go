package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/system-design-url-shortener/magikarp/entity"
	"go.uber.org/dig"
)

// Single tone logger for api.
var logger *log.Logger
var apiDevKey string

// Opts ...
type Opts struct {
	dig.In
	APIDevKey string `name:"api_dev_key"`
}

// RegisterAPI for API endpoints.
func RegisterAPI(newLogger *log.Logger, opts Opts, shortenURLService entity.ShortenURLService, store sessions.Store) (*gin.Engine, error) {

	logger = newLogger
	apiDevKey = opts.APIDevKey
	logger.Infof("RegisterAPI")

	r := gin.New()

	r.Use(sessions.Sessions("userSession", store))

	// Auth
	auth := r.Group("/auth")
	auth.Use(AuthenticationFunc())
	auth.GET("/ping", pingFunc())

	r.POST("/signup", SignupFunc(shortenURLService))
	r.POST("/login", LoginFunc(shortenURLService))
	r.GET("/logout", Logout)
	r.GET("/", pingFunc())

	// shorten url
	r.POST("/url", shortenerFunc(shortenURLService))

	// Issue about gin: https://github.com/gin-gonic/gin/issues/1730
	// Simply accept some flexibility lost now. Maybe need to change the api framework at the end.
	r.GET("/r/:shortURL", getFunc(shortenURLService))

	return r, nil
}
