package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/system-design-url-shortener/magikarp/entity"
	"net/http"
)

type shortenerFuncForm struct {
	APIDevKey   string    `form:"api_dev_key"`
	UserID      uuid.UUID `form:"userID"`
	OriginalURL string    `form:"originalURL" binding:"required"`
	ExpireTime  int64     `form:"expireTime"`
}

func shortenerFunc(service entity.ShortenURLService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form shortenerFuncForm
		if err := c.Bind(&form); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"msg": err.Error(),
				},
			)
			return
		}

		if form.APIDevKey != apiDevKey {
			logger.Debug("Not dev")
			// TODO: Check rate limit.
		}

		url, err := service.ShortenURL(
			form.OriginalURL,
			form.ExpireTime,
			form.UserID,
		)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"msg": err.Error(),
				},
			)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"url": url.ShortURL,
		})
	}
}
