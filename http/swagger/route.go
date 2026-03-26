package serverSwagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(router *gin.Engine, cfg DocsConfig) {
	if !cfg.Enabled {
		return
	}

	if cfg.Path == "" {
		cfg.Path = "/docs"
	}

	if cfg.SpecURL == "" {
		cfg.SpecURL = "/swagger/doc.json"
	}

	if cfg.Title == "" {
		cfg.Title = "API Docs"
	}

	router.Static("/docs/assets", "./http/swagger/assets")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET(cfg.Path, redocHandler(cfg))
}
