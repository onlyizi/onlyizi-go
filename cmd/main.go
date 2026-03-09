package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onlyizi/onlyizi-go/app"
	apperrors "github.com/onlyizi/onlyizi-go/errors"
	onlyiziHttp "github.com/onlyizi/onlyizi-go/http"
	"github.com/onlyizi/onlyizi-go/observability/logs"
	"github.com/onlyizi/onlyizi-go/observability/metrics"
)

func main() {
	logs.Init(logs.Config{
		Service:     "Onlyizi library",
		Environment: logs.Development,
		Version:     "0.1.0",
	})

	metrics.Init()
	metrics.InitHTTP("Onlyizi Library")

	httpServer := onlyiziHttp.NewServer(
		"Onlyizi Library server",
		":8080",
		testRoutes,
	)

	app.Run(httpServer)
}
func testRoutes(r *gin.Engine) {

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"test": "tested",
		})
	})

	r.GET("/error", func(ctx *gin.Context) {
		ctx.Error(apperrors.BadRequest(
			"invalid_input",
			"input is invalid",
		))
	})
}
