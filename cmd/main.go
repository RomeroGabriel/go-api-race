package main

import (
	moduleBrasilApi "RomeroGabriel/go-api-race/pkg/brasilapi"
	moduleViaCepApi "RomeroGabriel/go-api-race/pkg/viacepapi"
	"fmt"
)

func main() {
	brasilApi := moduleBrasilApi.NewBrasilApi()
	viaCepApi := moduleViaCepApi.NewViaCepApi()

	brasilApiCh := make(chan moduleBrasilApi.BrasilApiResult)
	viaCepApiCh := make(chan moduleViaCepApi.ViaCepResult)

	go func() {
		data, err := brasilApi.GetAddress("01153000")
		if err != nil {
			panic(err)
		}
		brasilApiCh <- *data
	}()
	go func() {
		data, err := viaCepApi.GetAddress("01153000")
		if err != nil {
			panic(err)
		}
		viaCepApiCh <- *data
	}()

	select {
	case address := <-brasilApiCh:
		fmt.Printf("BrasilApi: %s", address.ToString())
	case address := <-viaCepApiCh:
		fmt.Printf("ViaCep: %s", address.ToString())
	}
}
