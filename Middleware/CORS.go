package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORSMiddleware handles CORS requests
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, OPTIONS, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, Accept, Authorization, videourl, Access-Control-Allow-Headers, Access-Control-Allow-Origin, Access-Control-Request-Method, Access-Control-Request-Headers")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization, videourl")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
