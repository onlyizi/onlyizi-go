# Onlyizi Biblioteca

## Observabilidade

```
logs    -> promtail     -> loki
metrics -> prometheus
traces  -> jaeger
```

Grafana consulta os três:
    - loki
    - jaeger
    - prometheus