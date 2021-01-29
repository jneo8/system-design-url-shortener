package api

import (
	"github.com/gin-gonic/gin"
	"github.com/system-design-url-shortener/magikarp/entity"
	"net/http"
)

func getFunc(urlBackend entity.ShortenURLService) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortURL := c.Param("shortURL")

		url, err := urlBackend.GetByShortURL(shortURL)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"msg": err.Error(),
				},
			)
			return
		}
		// TODO: Log to analysis system.
		logger.Debug(c.Request, url)
		c.Redirect(http.StatusMovedPermanently, url)
	}
}
