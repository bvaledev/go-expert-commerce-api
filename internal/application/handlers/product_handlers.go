package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/bvaledev/go-expert-commerce-api/internal/domain/contracts"
	"github.com/bvaledev/go-expert-commerce-api/internal/domain/dto"
	"github.com/bvaledev/go-expert-commerce-api/internal/domain/entity"
	pkgEntity "github.com/bvaledev/go-expert-commerce-api/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB contracts.IProductRepository
}

func NewProductHandler(db contracts.IProductRepository) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create product godoc
//
//	@Summary		Create product
//	@Description	Create product
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.CreateProductDTO	true	"product request"
//	@Success		201
//	@Failure		500	{object}	error
//	@Router			/products [post]
//	@Security		ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDTO dto.CreateProductDTO
	err := json.NewDecoder(r.Body).Decode(&productDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := entity.NewProduct(productDTO.Name, productDTO.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Lists product godoc
//
//	@Summary		List products
//	@Description	List products
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			page	query		string	false	"page number"
//	@Param			limit	query		string	false	"page limit"
//	@Param			sort	query		string	false	"page sort order"
//	@Success		200		{array}		entity.Product
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/products [get]
//	@Security		ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("from handler", r.Context().Value("user"))
	pageQuery := r.URL.Query().Get("page")
	limitQuery := r.URL.Query().Get("limit")
	sortQuery := r.URL.Query().Get("sort")

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		limit = 10
	}

	products, err := h.ProductDB.FindAll(page, limit, sortQuery)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Get product godoc
//
//	@Summary		Get product
//	@Description	Get product
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"product id"	Format(uuid)
//	@Success		200	{array}		entity.Product
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/products/{id} [get]
//	@Security		ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	validId, err := pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Invalid ID"}`))
		return
	}

	product, err := h.ProductDB.FindByID(validId.String())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Update product godoc
//
//	@Summary		Update product
//	@Description	Update product
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"product id"	Format(uuid)
//	@Param			request	body		dto.UpdateProductDTO	true	"product request"
//	@Success		200		{array}		entity.Product
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Router			/products/{id} [put]
//	@Security		ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	validId, err := pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Invalid ID"}`))
		return
	}

	var productDTO dto.UpdateProductDTO
	err = json.NewDecoder(r.Body).Decode(&productDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Invalid Form"}`))
		return
	}

	product, err := h.ProductDB.FindByID(validId.String())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	product.Name = productDTO.Name
	product.Price = productDTO.Price

	err = h.ProductDB.Update(product)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Delete product godoc
//
//	@Summary		Delete product
//	@Description	Delete product
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"product id"	Format(uuid)
//	@Success		200	{array}		entity.Product
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/products/{id} [delete]
//	@Security		ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	validId, err := pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Invalid ID"}`))
		return
	}

	err = h.ProductDB.Delete(validId.String())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
