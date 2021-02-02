package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/system-design-url-shortener/magikarp/entity"
	"net/http"
)

// AuthenticationFunc ...
func AuthenticationFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("id")
		if sessionID == nil {
			c.JSON(http.StatusNotFound, gin.H{"meassage": "unauthorized"})
		}
		c.Next()
	}
}

// SignupFunc for sign up user.
func SignupFunc(shortenURLService entity.ShortenURLService) gin.HandlerFunc {

	return func(c *gin.Context) {
		var form SignupForm
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
		if !form.ValidPassword() {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"msg": "Password not equal",
				},
			)
			return
		}

		user, err := shortenURLService.NewUser(form.ToUser())
		if err != nil {
			logger.Debug(err)
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"msg": err.Error(),
				},
			)
			return
		}
		logger.Debug(user)
		c.JSON(
			http.StatusOK,
			gin.H{
				"msg": "Successfully Signup",
			},
		)
	}
}

// LoginFunc ...
func LoginFunc(shortenURLService entity.ShortenURLService) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userName := c.PostForm("userName")
		password := c.PostForm("password")
		if !shortenURLService.UserLogin(entity.User{UserName: userName, Password: password}) {
			c.JSON(http.StatusBadRequest, gin.H{"meassage": "Username or password incorrect"})
			return
		}
		session.Set("userName", userName)
		session.Save()
		logger.Infof("User login: %s", userName)
		c.JSON(http.StatusOK, gin.H{"meassage": "Sign In Successfully"})
	}
}

// Logout ...
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"meassage": "Sign out Successfully"})
}
