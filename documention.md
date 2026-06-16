# Dia 1 - Organizacao inicial do projeto EstateHub API

## 1. Objetivo do dia

No primeiro dia foi criada a base inicial da API do projeto **EstateHub**, usando Go e o pacote `net/http`.

O foco foi organizar o projeto em uma estrutura simples, separar responsabilidades e criar os primeiros endpoints de verificacao da aplicacao.

## 2. Modulo Go

O projeto foi iniciado como um modulo Go:

```go
module estatehub-api

go 1.26.1
```

Esse nome de modulo e usado nos imports internos do projeto, por exemplo:

```go
import "estatehub-api/internal/platform/config"
```

## 3. Estrutura de pastas criada

A estrutura atual do projeto esta organizada assim:

```text
EstateHub/
|-- cmd/
|   `-- api/
|       `-- main.go
|
|-- internal/
|   `-- platform/
|       |-- config/
|       |   `-- config.go
|       |
|       `-- http/
|           |-- health.go
|           |-- response.go
|           `-- route.go
|
|-- documention.md
`-- go.mod
```

## 4. Entrada da aplicacao

O arquivo principal da API fica em:

```text
cmd/api/main.go
```

Ele e responsavel por:

1. Carregar as configuracoes da aplicacao.
2. Criar o router HTTP.
3. Exibir no terminal o nome da aplicacao, ambiente e porta.
4. Iniciar o servidor HTTP.

Trecho principal:

```go
cfg := config.Load()

router := platformhttp.NewRouter()

log.Printf("%s running in %s mode on port %s\n", cfg.AppName, cfg.AppEnv, cfg.AppPort)

err := http.ListenAndServe(cfg.Addr(), router)
if err != nil {
	log.Fatal(err)
}
```

## 5. Configuracao da aplicacao

As configuracoes ficam em:

```text
internal/platform/config/config.go
```

Foi criada a struct `Config`:

```go
type Config struct {
	AppName string
	AppEnv  string
	AppPort string
}
```

A funcao `Load()` carrega valores das variaveis de ambiente e usa valores padrao quando elas nao existem:

```go
func Load() Config {
	return Config{
		AppName: getEnv("APP_NAME", "EstateHub API"),
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),
	}
}
```

Valores padrao:

```text
APP_NAME=EstateHub API
APP_ENV=development
APP_PORT=8080
```

Tambem foi criada a funcao `Addr()` para montar o endereco usado pelo servidor:

```go
func (c Config) Addr() string {
	return ":" + c.AppPort
}
```

## 6. Router HTTP

O router foi separado no arquivo:

```text
internal/platform/http/route.go
```

Foi criada a funcao `NewRouter()`, que monta um `ServeMux` e registra as rotas da API:

```go
func NewRouter() nethttp.Handler {
	mux := nethttp.NewServeMux()

	healthHandler := NewHealthHandler()

	mux.HandleFunc("/healtz", healthHandler.Healtz)
	mux.HandleFunc("/readz", healthHandler.Readz)

	return mux
}
```

Rotas criadas no Dia 1:

```text
GET /healtz
GET /readz
```

Observacao: a rota atual esta como `/healtz`. Se a ideia for seguir o padrao mais comum, futuramente pode ser ajustada para `/healthz`.

## 7. Health handler

Os handlers de saude da API ficam em:

```text
internal/platform/http/health.go
```

Foi criada a struct `HealthHandler`:

```go
type HealthHandler struct{}
```

E tambem a funcao construtora:

```go
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}
```

### Rota `/healtz`

Essa rota verifica se a API esta respondendo.

Ela aceita apenas o metodo `GET`.

Resposta de sucesso:

```json
{
  "status": "ok"
}
```

Caso o metodo nao seja `GET`, retorna erro `405 Method Not Allowed`.

### Rota `/readz`

Essa rota verifica se a API esta pronta para receber requisicoes.

Ela aceita apenas o metodo `GET`.

Resposta de sucesso:

```json
{
  "status": "ready"
}
```

Caso o metodo nao seja `GET`, retorna erro `405 Method Not Allowed`.

## 8. Helpers de resposta JSON

Os helpers para respostas HTTP em JSON ficam em:

```text
internal/platform/http/response.go
```

Foi criada a struct `ErrorResponse`:

```go
type ErrorResponse struct {
	Message string `json:"message"`
}
```

Tambem foram criadas duas funcoes auxiliares:

```go
func WriteJson(w nethttp.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("error encoding json:", err)
	}
}
```

```go
func ErrorJson(w nethttp.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(ErrorResponse{Message: message})
	if err != nil {
		log.Println("error encoding json:", err)
	}
}
```

Essas funcoes ajudam a manter o codigo dos handlers mais limpo, evitando repetir a configuracao de header e encoding JSON em cada rota.

## 9. Como rodar a API

Na raiz do projeto, executar:

```bash
go run ./cmd/api
```

Por padrao a API sobe na porta `8080`.

Exemplo de saida esperada no terminal:

```text
EstateHub API running in development mode on port 8080
```

## 10. Como testar as rotas

Testar health check:

```bash
curl http://localhost:8080/healtz
```

Resposta esperada:

```json
{
  "status": "ok"
}
```

Testar readiness:

```bash
curl http://localhost:8080/readz
```

Resposta esperada:

```json
{
  "status": "ready"
}
```

## 11. Resumo do que foi feito no Dia 1

1. Criado o modulo Go `estatehub-api`.
2. Criada a entrada da aplicacao em `cmd/api/main.go`.
3. Criada a camada de configuracao em `internal/platform/config`.
4. Criado o router HTTP em `internal/platform/http/route.go`.
5. Criado o handler de health check em `internal/platform/http/health.go`.
6. Criados helpers para respostas JSON em `internal/platform/http/response.go`.
7. Criadas as rotas iniciais `/healtz` e `/readz`.
8. Servidor HTTP configurado para rodar na porta `8080` por padrao.

## 12. Proximos passos sugeridos

1. Ajustar `/healtz` para `/healthz`, se quiser seguir o nome mais comum.
2. Criar uma camada de dominio para as entidades principais do EstateHub.
3. Criar os primeiros endpoints de imoveis.
4. Adicionar testes para os handlers HTTP.
5. Adicionar arquivo `.env` ou documentar melhor as variaveis de ambiente.


# Dia 2 - Organizacao inicial do projeto EstateHub API
## 1. Objetivo do dia 
No segundo dia, o foco foi criar a estrutura inicial do banco de dados para o projeto **EstateHub**.
Foram criadas as tabelas principais para representar os usuarios, agencias e corretores, usando SQL.
## 2. Script de migracao
O script de migracao para criar as tabelas foi criado no arquivo:
```textmigrations/001_init.sql
```
## 3. Tabela `users`
A tabela `users` foi criada para armazenar as informacoes basicas dos usuarios do sistema, como nome, email e telefone. Ela tem a seguinte estrutura:
```sqlCREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		phone VARCHAR(20) NOT NULL
	);
```
## 4- modelo de dados