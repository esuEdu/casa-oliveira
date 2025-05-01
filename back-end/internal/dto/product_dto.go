package dto

import "github.com/esuEdu/casa-oliveira/internal/entity"

type ProductPagination struct {
	Results  []entity.Product `json:"results"`
	Total    int64            `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"page_size"`
}

type UpdateProductInput struct {
	Name        *string   `json:"name"`
	Category    *string   `json:"category"`
	Price       *string   `json:"price"`
	Description *string   `json:"description"`
	ImageUrl    *[]string `json:"image_url"`
	// Add more fields as needed...
}
