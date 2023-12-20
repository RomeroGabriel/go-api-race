package viacepapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGetAddress(t *testing.T) {
	api := NewViaCepApi()
	address, err := api.GetAddress("01153000")
	assert.NoError(t, err)
	assert.NotNil(t, address)
	assert.Equal(t, "01153-000", address.Cep)
	assert.Equal(t, "SP", address.Uf)
	assert.Equal(t, "São Paulo", address.Localidade)
}
