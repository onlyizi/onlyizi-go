package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onlyizi/onlyizi-go/app"
	onlyiziHttp "github.com/onlyizi/onlyizi-go/http"
	"github.com/onlyizi/onlyizi-go/observability/logs"
)

func main() {

	logs.Init(logs.Config{
		Service:     "Onlyizi library",
		Environment: logs.Development,
		Version:     "0.1.0",
	})

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
}
