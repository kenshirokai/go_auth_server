package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORS(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.Writer.WriteHeader(http.StatusOK)
	}
	headers := c.Request.Header.Get("Access-Control-Request-Headers")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Request", "GET,POST,PUT,DELETE,OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", headers)
	c.Next()
}
