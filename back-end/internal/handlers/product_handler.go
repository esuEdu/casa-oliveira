package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/esuEdu/casa-oliveira/internal/dto"
	"github.com/esuEdu/casa-oliveira/internal/entity"
	"github.com/esuEdu/casa-oliveira/internal/service"
	"github.com/esuEdu/casa-oliveira/internal/util"
)

type ProductHandler struct {
	Service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{Service: s}
}

// Main method handler
func (h *ProductHandler) HandleProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		h.createProduct(w, r)
	case http.MethodGet:
		h.getProducts(w, r)
	default:
		util.RespondWithJSON(w, http.StatusMethodNotAllowed, util.APIResponse{
			Message: "Method Not Allowed",
			Error:   "Invalid_HTTP_Method",
		})
	}
}

func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		h.getProduct(w, r)
	case http.MethodPut:
		h.updateProduct(w, r)
	default:
		util.RespondWithJSON(w, http.StatusMethodNotAllowed, util.APIResponse{
			Message: "Method Not Allowed",
			Error:   "Invalid_HTTP_Method",
		})
	}
}

// POST /api/product
func (h *ProductHandler) createProduct(w http.ResponseWriter, r *http.Request) {

	// Get the product
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Request body require",
			Error:   "Body_missing",
		})
		return
	}

	// Validate required fields
	if product.Name == "" || product.Description == "" || product.Price == "" {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Missing or invalid fields",
			Error:   "invalid_product_data",
		})
		return
	}

	// Create product
	prod, err := h.Service.CreateProduct(&product)
	if err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Failed to create product",
			Error:   err.Error(),
		})
		return
	}

	// Response product
	util.RespondWithJSON(w, http.StatusCreated, util.APIResponse{
		Message: "Product created successfully",
		Data:    prod,
	})
}

// GET /api/product
func (h *ProductHandler) getProducts(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	// Default values
	page := 1
	pageSize := 10

	// Parse "page"
	if p := query.Get("page"); p != "" {
		if parsedPage, err := strconv.Atoi(p); err == nil {
			page = parsedPage
		}
	}

	// Parse "pageSize"
	if ps := query.Get("pageSize"); ps != "" {
		if parsedSize, err := strconv.Atoi(ps); err == nil {
			pageSize = parsedSize
		}
	}

	product, err := h.Service.ListProduct(page, pageSize)
	if err != nil {
		util.RespondWithJSON(w, http.StatusCreated, util.APIResponse{
			Message: "Failed to list products",
			Error:   err.Error(),
		})
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, util.APIResponse{
		Message: "Products list successfully",
		Data:    product,
	})
}

// GET /api/product/:id
func (h *ProductHandler) getProduct(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path // e.g., /api/product/123

	// Expecting: /api/product/:id
	parts := strings.Split(path, "/")
	if len(parts) != 4 || parts[1] != "api" || parts[2] != "product" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id := parts[3]

	product, err := h.Service.GetProduct(id)
	if err != nil {
		util.RespondWithJSON(w, http.StatusCreated, util.APIResponse{
			Message: "Failed to get product",
			Error:   err.Error(),
		})
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, util.APIResponse{
		Message: "Products get successfully",
		Data:    product,
	})

}

func (h *ProductHandler) updateProduct(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path // e.g., /api/product/123

	// Expecting: /api/product/:id
	parts := strings.Split(path, "/")
	if len(parts) != 4 || parts[1] != "api" || parts[2] != "product" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id := parts[3]

	// Get the product
	var input dto.UpdateProductInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Request body require",
			Error:   "Body_missing",
		})
		return
	}

	product, err := h.Service.UpdateProduct(id, &input)
	if err != nil {
		util.RespondWithJSON(w, http.StatusCreated, util.APIResponse{
			Message: "Failed to update product",
			Error:   err.Error(),
		})
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, util.APIResponse{
		Message: "Product update successfully",
		Data:    product,
	})

}
