package middlewares

import (
	"net/http"
	"ai-saas-schedular-server/server/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		claims, msg := helpers.ValidateToken(clientToken)
		if msg != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)

		c.Next()
	}
}
