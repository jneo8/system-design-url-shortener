package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/system-design-url-shortener/kgs/entity"
	"net/http"
	"strconv"
)

func getFunc(logger *log.Logger, repo entity.KeyRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get params
		var expire int64
		if v, err := strconv.ParseInt(c.DefaultQuery("expire", "0"), 10, 64); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"message": err,
				},
			)
		} else {
			expire = v
		}

		// Get key
		key, err := repo.GetKey(expire)
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
