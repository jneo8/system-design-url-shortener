package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
	"github.com/system-design-url-shortener/magikarp/entity"
	"net/http"
)

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
			logger.Debug("Not dev")
			// TODO: Check rate limit.
		}

		session := sessions.Default(c)
		userName := session.Get("userName")
		logger.Debug(userName)

		url, err := service.ShortenURL(
			form.OriginalURL,
			form.ExpireTime,
			nil,
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
			"url": schema + "://" + c.Request.Host + "/r/" + url.ShortURL,
		})
	}
}
