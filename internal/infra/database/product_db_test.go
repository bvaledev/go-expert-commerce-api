package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bvaledev/go-expert-commerce-api/internal/entity"
	"github.com/bvaledev/go-expert-commerce-api/pkg/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestProduct_ShouldCreateProductWithCorrectValues(t *testing.T) {
	db := testhelper.SetupDBTest(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 18.982)
	assert.NoError(t, err)

	productDB := NewProductDB(db)

	err = productDB.Create(product)
	assert.NoError(t, err)

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID.String()).Error
	assert.NoError(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, "Product 1", productFound.Name)
	assert.Equal(t, 18.98, productFound.Price)
}

func TestProduct_ShouldFindPaginatedProductList(t *testing.T) {
	db := testhelper.SetupDBTest(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %v", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}
	productDB := NewProductDB(db)

	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestProduct_ShouldFindProductById(t *testing.T) {
	db := testhelper.SetupDBTest(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 18.982)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProductDB(db)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, "Product 1", productFound.Name)
	assert.Equal(t, 18.98, productFound.Price)
}

func TestProduct_ShouldUpdateProduct(t *testing.T) {
	db := testhelper.SetupDBTest(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 18.982)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProductDB(db)

	product.Name = "Updated Product"
	product.Price = 18.59
	productDB.Update(product)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, "Updated Product", productFound.Name)
	assert.Equal(t, 18.59, productFound.Price)
}

func TestProduct_ShouldDeleteProduct(t *testing.T) {
	db := testhelper.SetupDBTest(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 18.982)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProductDB(db)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, productFound)
}
