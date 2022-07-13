package application

import (
	"github.com/codeedu/go-hexagonal/application"
	"testing"
	"github.com/stretch/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	Product.Price = 10

	err := product.Enable()
	require.Nil(t, err)
}