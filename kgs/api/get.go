package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/system-design-url-shortener/kgs/entity"
	"net/http"
)

func getFunc(logger *log.Logger, repo entity.KeyRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		key, err := repo.GetKey()
		if err != nil {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"message": "No key been found.",
				},
			)
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": key,
			},
		)
	}
}
