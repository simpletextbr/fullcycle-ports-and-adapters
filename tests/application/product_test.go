package application_test

import (
	"testing"

	"github.com/simpletextbr/fullcycle-ports-and-adapters/application"
	"github.com/stretchr/testify/require"
)

func TestApplicationProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}
