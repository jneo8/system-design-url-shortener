package api

import (
	"github.com/gin-gonic/gin"
	"github.com/system-design-url-shortener/magikarp/entity"
	"net/http"
)

func getFunc(service entity.ShortenURLService) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortURL := c.Param("shortURL")
		url, err := service.GetByShortURL(shortURL)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"msg": err.Error(),
				},
			)
			return
		}
		c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
	}
}
