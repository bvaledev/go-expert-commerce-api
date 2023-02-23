package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct_ShouldCreateNewProductWithCorrectValues(t *testing.T) {
	product, err := NewProduct("Produto 1", 3567.9592)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID.String())
	assert.Equal(t, "Produto 1", product.Name)
	assert.Equal(t, 3567.96, product.Price)
}

func TestProduct_ShouldFormatProductPrice(t *testing.T) {
	product, err := NewProduct("Produto 1", 3567.9592)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, 3567.96, product.Price)

	product, err = NewProduct("Produto 1", 3567.9500)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, 3567.95, product.Price)

	product, err = NewProduct("Produto 1", 3567.9544)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, 3567.95, product.Price)
}

func TestProduct_ShouldReturnErrorNameIsRequiredValues(t *testing.T) {
	product, err := NewProduct("", 3567.9592)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProduct_ShouldReturnErrorPriceIsRequiredValues(t *testing.T) {
	product, err := NewProduct("Produto 1", 0)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProduct_ShouldReturnErrorPriceIsInvalidValues(t *testing.T) {
	product, err := NewProduct("Produto 1", -58.5)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, ErrPriceIsInvalid, err)
}
