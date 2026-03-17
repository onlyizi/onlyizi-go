package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MetricsIPAllowlist(allowedIPs []string) gin.HandlerFunc {
	allowed := make(map[string]struct{}, len(allowedIPs))
	for _, ip := range allowedIPs {
		allowed[ip] = struct{}{}
	}

	return func(c *gin.Context) {

		ip := c.ClientIP()

		if _, ok := allowed[ip]; !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
