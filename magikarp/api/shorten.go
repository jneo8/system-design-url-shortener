package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/system-design-url-shortener/magikarp/entity"
	"net/http"
)

type shortenerFuncForm struct {
	APIDevKey   string `form:"api_dev_key"`
	UserID      string `form:"userID"`
	OriginalURL string `form:"originalURL" binding:"required"`
	ExpireTime  int64  `form:"expireTime"`
}

func shortenerFunc(service entity.ShortenURLService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form shortenerFuncForm
		if err := c.Bind(&form); err != nil {
			logger.Debug(err)
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"msg": err.Error(),
				},
			)
			return
		}

		if form.APIDevKey != apiDevKey {
			// TODO: Check rate limit.
		}

		var userID *uuid.UUID
		if userUUID, err := uuid.Parse(form.UserID); err == nil {
			userID = &userUUID
		}

		url, err := service.ShortenURL(
			form.OriginalURL,
			form.ExpireTime,
			userID,
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

		var schema string
		switch c.Request.Proto {
		case "HTTP/1.1":
			schema = "http"
		case "HTTPS/2":
			schema = "https"
		default:
			schema = "http"
		}

		c.JSON(http.StatusOK, gin.H{
			"url": schema + "://" + c.Request.Host + "/" + url.ShortURL,
		})
	}
}
