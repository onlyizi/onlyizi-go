package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/onlyizi/onlyizi-go/app"
	"github.com/onlyizi/onlyizi-go/config"
	"github.com/onlyizi/onlyizi-go/errors"
	onlyiziHttp "github.com/onlyizi/onlyizi-go/http"
	"github.com/onlyizi/onlyizi-go/http/middlewares"
	"github.com/onlyizi/onlyizi-go/infra/postgres"
	"github.com/onlyizi/onlyizi-go/infra/redis"
	"github.com/onlyizi/onlyizi-go/observability"
)

/*
Example Server

Este arquivo demonstra como iniciar um serviço utilizando a biblioteca onlyizi-go.

Ele serve para três propósitos:

1. Testar manualmente a biblioteca durante o desenvolvimento.
2. Servir como documentação de referência para novos serviços Onlyizi.
3. Mostrar o fluxo completo de inicialização de um serviço.

Fluxo de inicialização:

1. Inicializar observabilidade (logs, metrics, tracing)
2. Criar servidor HTTP
3. Registrar rotas
4. Iniciar runtime de serviços
*/

func main() {
	// --------------------------------------------------
	// Inicializa variáveis de ambiente
	// --------------------------------------------------
	config.LoadEnv(".env.example")

	service := config.ServiceConfig()
	httpCfg := config.HTTPConfig()

	// --------------------------------------------------
	// Inicializa observabilidade
	// --------------------------------------------------
	err := observability.Init(observability.Config{
		ServiceName: service.Name,
		Version:     service.Version,
		Environment: service.Environment,
	})
	if err != nil {
		panic(err)
	}

	// --------------------------------------------------
	// Cria servidor HTTP
	// --------------------------------------------------
	cors := middlewares.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
	}
	httpServer := onlyiziHttp.NewServer(
		service.Name+"-http",
		":"+strconv.Itoa(httpCfg.Port),
		cors,
		registerRoutes,
	)

	// --------------------------------------------------
	// Executa serviços
	// --------------------------------------------------
	app.Run(
		postgres.New(),
		redis.New(),
		httpServer,
	)
}

/*
registerRoutes

Função responsável por registrar as rotas HTTP do serviço.

Cada aplicação que utiliza onlyizi-go deve implementar
uma função similar para registrar suas rotas.
*/
func registerRoutes(r *gin.Engine) {

	r.GET("/example", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"message": "example endpoint working",
		})
	})

	r.GET("/error", func(ctx *gin.Context) {

		ctx.Error(errors.BadRequest(
			"example_error",
			"this is an example error",
		))
	})
}
