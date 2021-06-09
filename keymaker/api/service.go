package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/system-design-url-shortener/keymaker/entity"
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
