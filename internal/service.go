package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	API          string
	CEP          string
	Street       string
	Neighborhood string
	City         string
	State        string
}

type BrasilAPIResponse struct {
	CEP          string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type ViaCEPResponse struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

func FetchBrasilAPI(ctx context.Context, cep string, ch chan<- Result) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var data BrasilAPIResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return
	}

	select {
	case ch <- Result{
		API:          "BrasilAPI",
		CEP:          data.CEP,
		Street:       data.Street,
		Neighborhood: data.Neighborhood,
		City:         data.City,
		State:        data.State,
	}:
	case <-ctx.Done():
	}
}

func FetchViaCEP(ctx context.Context, cep string, ch chan<- Result) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var data ViaCEPResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return
	}

	select {
	case ch <- Result{
		API:          "ViaCEP",
		CEP:          data.CEP,
		Street:       data.Logradouro,
		Neighborhood: data.Bairro,
		City:         data.Localidade,
		State:        data.UF,
	}:
	case <-ctx.Done():
	}
}
