package viacepapi

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type ViaCepApi struct {
	baseurl_api string
}

type ViaCepResult struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewViaCepApi() *ViaCepApi {
	return &ViaCepApi{
		baseurl_api: "http://viacep.com.br/ws/",
	}
}

func (api ViaCepApi) GetAddress(code string) (*ViaCepResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	url := api.baseurl_api + code + "/json"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Request Timeout consulting the " + api.baseurl_api + " API")
		return nil, context.Canceled
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result ViaCepResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		log.Println("Request Timeout consulting the " + api.baseurl_api + " API")
		return nil, context.Canceled
	default:
		return &result, nil
	}
}
