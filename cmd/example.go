package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/onlyizi/onlyizi-go/app"
	"github.com/onlyizi/onlyizi-go/bootstrap"
	"github.com/onlyizi/onlyizi-go/config"
	"github.com/onlyizi/onlyizi-go/errors"
	onlyiziHttp "github.com/onlyizi/onlyizi-go/http"
	"github.com/onlyizi/onlyizi-go/http/middlewares"
	serverSwagger "github.com/onlyizi/onlyizi-go/http/swagger"
	"github.com/onlyizi/onlyizi-go/infra/postgres"
	"github.com/onlyizi/onlyizi-go/infra/redis"
	"github.com/onlyizi/onlyizi-go/observability"

	_ "github.com/onlyizi/onlyizi-go/docs"
)

/*
Example Server

Este arquivo demonstra a forma recomendada de inicializar um serviço HTTP
utilizando a biblioteca onlyizi-go após a introdução do bootstrap.Start.

Objetivos deste exemplo:

1. Servir como referência oficial de inicialização para novos serviços.
2. Demonstrar o fluxo padronizado de bootstrap da biblioteca.
3. Mostrar a separação entre:
   - configuração do ambiente
   - observabilidade
   - serviços de infraestrutura
   - servidor HTTP
   - registro de rotas

Ideia central do bootstrap.Start:

A função bootstrap.Start centraliza o fluxo padrão de inicialização de uma aplicação.
Com isso, evitamos repetir em cada serviço a mesma sequência de passos, como:

- carregar variáveis de ambiente
- inicializar observabilidade
- iniciar serviços de infraestrutura
- criar e iniciar servidor HTTP
- controlar shutdown da aplicação

Esse padrão reduz duplicação, melhora consistência entre serviços
e evita problemas de ordem de inicialização.

Fluxo executado por bootstrap.Start:

1. Carrega variáveis de ambiente, se um arquivo for informado
2. Inicializa observabilidade
3. Inicia serviços de bootstrap de forma sequencial
   Ex.: postgres, redis
4. Cria o servidor HTTP, se configurado
5. Inicia os serviços de runtime
   Ex.: servidor HTTP, workers, consumers
6. Aguarda sinal do sistema para shutdown
7. Finaliza todos os serviços em ordem reversa

Importante:

Este exemplo continua mostrando explicitamente as rotas e os serviços
utilizados, mas transfere para a biblioteca o controle do lifecycle da aplicação.
*/

// @Title Onlyizi Example Api
// @Description Essa é uma api de exemplo
func main() {
	/*
		Lemos as configurações básicas do serviço e do HTTP.

		Aqui estamos usando os helpers do pacote config para obter:
		- nome do serviço
		- versão
		- ambiente
		- porta HTTP

		Observação importante:
		Neste exemplo, config.LoadEnv(".env.example") não é mais chamado manualmente.
		Quem passa a fazer isso é o próprio bootstrap.Start, por meio do campo EnvFile.
	*/
	service := config.ServiceConfig()
	httpCfg := config.HTTPConfig()

	/*
		Configuração padrão de CORS para o servidor HTTP do exemplo.

		Esse bloco representa a configuração de infraestrutura HTTP da aplicação,
		e não a lógica de negócio em si.
	*/
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

	/*
		Chamamos bootstrap.Start para delegar à biblioteca o fluxo padronizado
		de inicialização da aplicação.

		O que estamos informando aqui:

		- EnvFile:
		  arquivo de ambiente que deve ser carregado antes de tudo

		- Observability:
		  configuração usada para inicializar logs, métricas e tracing

		- Bootstrap:
		  serviços que precisam iniciar antes do runtime
		  Ex.: conexões com banco, redis, infraestrutura crítica

		- HTTP:
		  configuração do servidor HTTP que a biblioteca deve criar

		Nesse modelo, postgres e redis sobem primeiro.
		Só depois o servidor HTTP é criado e iniciado.

		Isso evita problemas em que rotas, middlewares ou handlers tentam usar
		dependências que ainda não foram conectadas.
	*/
	err := bootstrap.Start(bootstrap.Config{
		EnvFile: ".env.example",

		Observability: observability.Config{
			ServiceName: service.Name,
			Version:     service.Version,
			Environment: service.Environment,
		},

		Bootstrap: []app.Service{
			postgres.New(),
			redis.New(),
		},

		HTTP: &bootstrap.HTTPConfig{
			Name: service.Name + "-http",
			Addr: ":" + strconv.Itoa(httpCfg.Port),
			CORS: cors,
			Docs: &serverSwagger.DocsConfig{
				Enabled: true,
				Title:   "Onlyizi library",
				Path:    "/docs",
				Product: "Library API",
			},
			Routes: []onlyiziHttp.RegisterRoutes{
				registerRoutes,
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

/*
registerRoutes

Função responsável por registrar as rotas HTTP da aplicação.

A biblioteca onlyizi-go cuida do bootstrap, do lifecycle e da infraestrutura
do servidor, mas cada aplicação continua responsável por definir suas próprias rotas.

Isso mantém a separação correta de responsabilidades:

- a biblioteca fornece a base de infraestrutura
- a aplicação fornece a lógica e os endpoints do domínio

Neste exemplo, temos:
- uma rota simples de sucesso
- uma rota que demonstra o fluxo padronizado de erro HTTP
*/
func registerRoutes(r *gin.Engine) {
	/*
		Endpoint simples para validar que o servidor está funcionando
		e que as rotas customizadas foram registradas corretamente.
	*/
	r.GET("/example", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "example endpoint working",
		})
	})

	/*
		Endpoint de exemplo para demonstrar o uso do pacote de erros padronizados.

		A middleware de tratamento de erros da biblioteca será responsável
		por transformar esse erro em uma resposta HTTP consistente.
	*/
	r.GET("/error", func(ctx *gin.Context) {
		ctx.Error(errors.BadRequest(
			"example_error",
			"this is an example error",
		))
	})
}
