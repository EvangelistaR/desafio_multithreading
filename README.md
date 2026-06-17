# Multithreading CEP Challenge

Aplicação em Go que consulta um CEP simultaneamente em duas APIs:

- BrasilAPI
- ViaCEP

A primeira API que responder é utilizada e a outra é descartada.

## Tecnologias

- Golang
- Goroutines
- Channels
- Select
- Context
- net/http

## Executando

Clone o projeto:

```bash
git clone <repo>
cd <repo>
```

Execute:

```bash
go run main.go
```

## Alterando o CEP

Passe um CEP específico no final do comando:

```bash
go run main.go 01310200
```

## Timeout

O timeout é de 1 segundo:

```go
context.WithTimeout(context.Background(), time.Second)
```

Caso nenhuma API responda dentro do prazo:

```bash
Erro: timeout ao consultar CEP
```

## Conceitos aplicados

- Concorrência com Goroutines
- Comunicação via Channels
- Race entre APIs
- Cancelamento com Context
- Timeout de requisição