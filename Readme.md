# Versionamento da Biblioteca
Este repositório segue **Semantic Versioning (SemVer)** para controle de versões.
As versões da biblioteca são publicadas através de **tags no GitHub** criadas durante o merge da branch `develop` para `master`.
A branch `master` representa **versões estáveis da biblioteca**, enquanto a branch `develop` é utilizada para desenvolvimento contínuo.

---
# Estrutura de Branches
O fluxo de desenvolvimento segue o modelo abaixo:

feature → develop → master → tag

### Branches
**develop**

* Branch principal de desenvolvimento
* Recebe features, melhorias e correções
* Pode conter código ainda não versionado

**master**

* Contém apenas versões estáveis da biblioteca
* Não recebe commits diretos
* Apenas merges da branch `develop`

---
# Processo de Release

Quando a biblioteca estiver pronta para uma nova versão:

1. Abrir um **Pull Request de `develop` para `master`**
2. Após aprovação, realizar o **merge**
3. Criar uma **tag de versão no GitHub** apontando para o commit em `master`

Exemplo de versões:

```
v0.1.0
v0.2.0
v0.3.0
```

Após criar a tag, a nova versão da biblioteca ficará disponível para consumo pelos projetos.

---

# Semantic Versioning

O versionamento segue o padrão:

```
vMAJOR.MINOR.PATCH
```

### MAJOR

Mudanças que quebram compatibilidade com versões anteriores.

Exemplo:

```
v2.0.0
```

### MINOR

Novas funcionalidades compatíveis com versões anteriores.

Exemplo:

```
v1.3.0
```

### PATCH

Correções de bugs ou pequenas melhorias sem alterar a API.

Exemplo:

```
v1.2.1
```

---

# Utilização nos Projetos

Para adicionar a biblioteca a um projeto Go:

```
go get github.com/onlyizi/onlyizi-go@vX.X.X
```

Exemplo:

```
go get github.com/onlyizi/onlyizi-go@v0.1.0
```

Para atualizar a dependência:

```
go get -u github.com/onlyizi/onlyizi-go
```

---

# Observações
* Nunca realizar commits diretos na branch `master`
* Toda evolução da biblioteca deve ocorrer através da branch `develop`
* Cada merge em `master` deve resultar em uma **tag de versão**


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

## Swagger

Rotas disponibilizados:
`GET /swagger/index.html`
`GET /docs   (ou o Path que você definir)`