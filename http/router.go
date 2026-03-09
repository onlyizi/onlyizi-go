package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onlyizi/onlyizi-go/http/middlewares"
)

type RegisterRoutes func(*gin.Engine)

func NewRouter(routes ...RegisterRoutes) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(GinMiddleware(middlewares.ObservabilityMiddleware))
	router.Use(middlewares.ErrorHandlerMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	for _, register := range routes {
		register(router)
	}

	return router
}
