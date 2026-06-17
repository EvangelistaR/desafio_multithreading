package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/EvangelistaR/desafio_multithreading/internal"
)

func main() {
	cep := "01153000"

	if len(os.Args) > 1 {
		cep = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resultChan := make(chan internal.Result)

	go internal.FetchBrasilAPI(ctx, cep, resultChan)
	go internal.FetchViaCEP(ctx, cep, resultChan)

	select {
	case result := <-resultChan:
		fmt.Printf("API vencedora: %s\n\n", result.API)
		fmt.Printf("CEP: %s\n", result.CEP)
		fmt.Printf("Logradouro: %s\n", result.Street)
		fmt.Printf("Bairro: %s\n", result.Neighborhood)
		fmt.Printf("Cidade: %s\n", result.City)
		fmt.Printf("Estado: %s\n", result.State)

	case <-ctx.Done():
		fmt.Println("Erro: timeout ao consultar CEP")
	}
}
