package brasilapi

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type BrasilApi struct {
	baseurl_api string
}

type BrasilApiResult struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func NewBrasilApi() *BrasilApi {
	return &BrasilApi{
		baseurl_api: "https://brasilapi.com.br/api/cep/v1/",
	}
}

func (api BrasilApi) GetAddress(code string) (*BrasilApiResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	url := api.baseurl_api + code
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
	var result BrasilApiResult
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
