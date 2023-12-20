package brasilapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGetAddress(t *testing.T) {
	api := NewBrasilApi()
	address, err := api.GetAddress("01153000")
	assert.NoError(t, err)
	assert.NotNil(t, address)
	assert.Equal(t, "01153000", address.Cep)
	assert.Equal(t, "SP", address.State)
	assert.Equal(t, "São Paulo", address.City)
}
