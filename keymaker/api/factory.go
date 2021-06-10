package api

import (
	"jneo8/system-design-url-shortener/keymaker/entity"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// New factory func.
func New(
	logger *log.Logger,
	repo entity.KeyRepository,
	engine *gin.Engine,
) (entity.APIService, error) {
	s := &service{Logger: logger, Repo: repo, Engine: engine}
	return s, nil
}

// Register gin Engine
func Register(logger *log.Logger, repo entity.KeyRepository) (*gin.Engine, error) {
	r := gin.New()
	r.GET("/ping", pingFunc())
	r.GET("/newKey", getFunc(logger, repo))
	return r, nil
}
