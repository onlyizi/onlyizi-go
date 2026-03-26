package bootstrap

import (
	"github.com/onlyizi/onlyizi-go/app"
	"github.com/onlyizi/onlyizi-go/config"
	onlyiziHttp "github.com/onlyizi/onlyizi-go/http"
	"github.com/onlyizi/onlyizi-go/http/middlewares"
	serverSwagger "github.com/onlyizi/onlyizi-go/http/swagger"
	"github.com/onlyizi/onlyizi-go/observability"
)

type HTTPConfig struct {
	Name   string
	Addr   string
	CORS   middlewares.CORSConfig
	Routes []onlyiziHttp.RegisterRoutes
	Docs   *serverSwagger.DocsConfig
}

type Config struct {
	EnvFile       string
	Observability observability.Config
	Bootstrap     []app.Service
	Runtime       []app.Service
	HTTP          *HTTPConfig
}

func Start(cfg Config) error {
	if cfg.EnvFile != "" {
		config.LoadEnv(cfg.EnvFile)
	}

	if err := observability.Init(cfg.Observability); err != nil {
		return err
	}

	runtime := make([]app.Service, 0, len(cfg.Runtime)+1)
	runtime = append(runtime, cfg.Runtime...)

	if cfg.HTTP != nil {
		httpServer := onlyiziHttp.NewServer(
			cfg.HTTP.Name,
			cfg.HTTP.Addr,
			cfg.HTTP.CORS,
			cfg.HTTP.Routes...,
		)

		if cfg.HTTP.Docs != nil && cfg.HTTP.Docs.Enabled {
			httpServer.WithDocs(onlyiziHttp.DocsConfig{
				Title: cfg.HTTP.Docs.Title,
				Path:  cfg.HTTP.Docs.Path,
			})
		}

		runtime = append(runtime, httpServer)
	}

	return app.Run(cfg.Bootstrap, runtime)
}
