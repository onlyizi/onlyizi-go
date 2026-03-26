package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onlyizi/onlyizi-go/http/middlewares"
	serverSwagger "github.com/onlyizi/onlyizi-go/http/swagger"
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

	// Health godoc
	// @Summary Return health of api
	// @Router /health [get]
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Metrics godoc
	// @Summary Return metrics of api
	// @Router /metrics [get]
	router.GET(
		"/metrics",
		middlewares.MetricsIPAllowlist([]string{"127.0.0.1", "::1"}),
		gin.WrapH(metrics.Handler()),
	)

	serverSwagger.Setup(router, serverSwagger.Config{
		Title: "Example Api",
	})
}
