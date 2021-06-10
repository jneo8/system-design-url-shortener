package api

import (
	"net/http"

	"jneo8/system-design-url-shortener/keymaker/entity"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func getFunc(logger *log.Logger, repo entity.KeyRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get key
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
