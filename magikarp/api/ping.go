package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func pingFunc() gin.HandlerFunc { // nolint
	return func(c *gin.Context) {
		logger.Info("ping")
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "pong",
			},
		)
	}
}
