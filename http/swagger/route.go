package serverSwagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	Title   string
	SpecURL string
	Path    string
}

func Setup(router *gin.Engine, cfg Config) {
	if cfg.Path == "" {
		cfg.Path = "/docs"
	}

	if cfg.SpecURL == "" {
		cfg.SpecURL = "/swagger/doc.json"
	}

	if cfg.Title == "" {
		cfg.Title = "API Docs"
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET(cfg.Path, redocHandler(cfg))
}
