package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinMiddleware(mw func(http.Handler) http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Next()
		}))

		handler.ServeHTTP(c.Writer, c.Request)
	}
}
