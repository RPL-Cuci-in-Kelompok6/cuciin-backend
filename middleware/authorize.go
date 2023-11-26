package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize(role int) func(*gin.Context) {
	return func(c *gin.Context) {
		userRole := c.GetInt("role")
		if userRole != role {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
