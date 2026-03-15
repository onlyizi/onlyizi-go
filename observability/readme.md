# Observabilidade

Este documento descreve como funciona a **observabilidade padrão das aplicações OnlyIZI**.

A observabilidade é baseada em três pilares fundamentais:

- **Logs**
- **Métricas**
- **Tracing**

A biblioteca `onlyizi-go/observability` padroniza a inicialização e o uso desses três componentes para todos os serviços.

O objetivo é garantir que **todas as aplicações tenham observabilidade consistente**, sem necessidade de configuração manual em cada serviço.

---

# Arquitetura Geral

O fluxo de observabilidade das aplicações funciona da seguinte forma:

```
applications
↓
logs → Loki
metrics → Prometheus
traces → Jaeger
↓
Grafana (visualização)
```

Cada aplicação produz logs, métricas e traces utilizando a biblioteca `observability`.

Esses dados são coletados pela infraestrutura de observabilidade e visualizados através do **Grafana**.

---

# Inicialização da Observabilidade

A observabilidade deve ser inicializada no início da aplicação:

```go
observability.Init(observability.Config{
    ServiceName: service.Name,
    Version:     service.Version,
    Environment: service.Environment,
})
```

Essa inicialização configura automaticamente:

logger estruturado

métricas OpenTelemetry

tracing distribuído

Logs

Os logs utilizam o logger estruturado baseado em zap.

Todos os logs seguem o formato JSON estruturado, contendo campos padrão:

service

environment

version

timestamp

level

message

Exemplo:

{
  "timestamp": "2026-03-15T20:00:00Z",
  "level": "info",
  "service": "users-api",
  "environment": "development",
  "version": "1.0.0",
  "message": "user created"
}
Fluxo de Logs

Os logs seguem o fluxo abaixo:

1. Aplicação (zap → stdout)
        ↓
2. Docker logs
        ↓
3. Promtail
        ↓
4. Loki
        ↓
5. Grafana

As aplicações escrevem logs no stdout

Docker coleta os logs dos containers

Promtail envia os logs para o Loki

Loki armazena os logs

Grafana permite buscar e visualizar logs

Métricas

As métricas utilizam OpenTelemetry Metrics.

As aplicações expõem métricas no endpoint:

/metrics

Esse endpoint é coletado pelo Prometheus.

Inicialização das Métricas

A biblioteca inicializa automaticamente o provider de métricas:

metrics.Init()

Para métricas HTTP:

metrics.InitHTTP(serviceName)
Métricas HTTP disponíveis

A biblioteca já inclui métricas padrão para APIs HTTP:

Total de requisições
http_requests_total

Contador de requisições HTTP.

Labels:

method

path

status

Latência das requisições
http_request_duration_ms

Histograma da duração das requisições HTTP em milissegundos.

Labels:

method

path

status

Fluxo de Métricas
application
    ↓
OpenTelemetry metrics
    ↓
/metrics endpoint
    ↓
Prometheus
    ↓
Grafana

Prometheus coleta métricas periodicamente e Grafana permite criar dashboards.

Tracing

Tracing distribuído permite acompanhar o caminho de uma requisição através de diferentes serviços.

A biblioteca utiliza OpenTelemetry Tracing.

Middleware HTTP

Para rastrear requisições HTTP, utilize o middleware:

tracing.Middleware

Cada requisição gera automaticamente um span.

Exemplo de nome do span:

GET /users
POST /orders
Inicialização do Tracing
tracing.Init(serviceName)

Isso configura o tracer provider da aplicação.

Fluxo de Tracing
application
    ↓
OpenTelemetry tracer
    ↓
Jaeger
    ↓
Jaeger UI / Grafana

Cada requisição gera um trace contendo:

duração da requisição

spans internos

erros

dependências entre serviços

Contexto de Logger

O logger pode ser propagado via context.Context.

Adicionar logger ao contexto:

ctx = logs.WithLogger(ctx, logger)

Recuperar logger:

logger := logs.FromContext(ctx)

Isso permite que diferentes partes da aplicação compartilhem o mesmo logger.

Shutdown

Durante o encerramento da aplicação, é recomendado finalizar a observabilidade:

observability.Shutdown(ctx)

Isso garante que:

spans pendentes sejam exportados

recursos do tracer sejam liberados