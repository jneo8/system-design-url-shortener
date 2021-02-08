package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/system-design-url-shortener/kgs/entity"
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
