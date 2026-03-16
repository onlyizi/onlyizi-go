package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onlyizi/onlyizi-go/http/middlewares"
	"github.com/onlyizi/onlyizi-go/observability/metrics"
)

type RegisterRoutes func(*gin.Engine)

func NewRouter(cors middlewares.CORSConfig, routes ...RegisterRoutes) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	router.Use(GinMiddleware(middlewares.CORSMiddleware(cors)))
	router.Use(GinMiddleware(middlewares.ObservabilityMiddleware))
	router.Use(middlewares.ErrorHandlerMiddleware())

	standardRoutes(router)

	for _, register := range routes {
		register(router)
	}

	return router
}

func standardRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/metrics", gin.WrapH(metrics.Handler()))
}