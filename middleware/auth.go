package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const AuthToken = "static-token"

func Authenticate(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != AuthToken {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token Not Valid",
		})
		return
	}
}
