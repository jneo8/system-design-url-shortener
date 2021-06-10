package api

import (
	"jneo8/system-design-url-shortener/keymaker/entity"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type service struct {
	Logger *log.Logger
	Engine *gin.Engine
	Repo   entity.KeyRepository
}

func (s service) Run() error {
	// Register
	if err := s.Engine.Run(); err != nil {
		return err
	}
	return nil
}
