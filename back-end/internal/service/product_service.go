package service

import (
	"strconv"

	"github.com/esuEdu/casa-oliveira/internal/dto"
	"github.com/esuEdu/casa-oliveira/internal/entity"
	"github.com/esuEdu/casa-oliveira/internal/repositories"
)

type ProductService interface {
	CreateProduct(p *entity.Product) (*entity.Product, error)
	GetProduct(id string) (*entity.Product, error)
	ListProduct(page, pageSize int) (*dto.ProductPagination, error)
	UpdateProduct(idStr string, input *dto.UpdateProductInput) (*entity.Product, error)
}

type productService struct {
	repo repositories.ProductRepo
}

func NewProductService(r repositories.ProductRepo) ProductService {
	return &productService{repo: r}
}

func (s *productService) CreateProduct(p *entity.Product) (*entity.Product, error) {

	err := s.repo.Create(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *productService) ListProduct(page, pageSize int) (*dto.ProductPagination, error) {

	productPagination, err := s.repo.List(page, pageSize)
	if err != nil {
		return nil, err
	}

	return productPagination, nil
}

func (s *productService) GetProduct(idStr string) (*entity.Product, error) {

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return nil, err
	}

	product, err := s.repo.FindByID(uint(id))
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) UpdateProduct(idStr string, input *dto.UpdateProductInput) (*entity.Product, error) {

	// product without update
	product, err := s.GetProduct(idStr)
	if err != nil {
		return nil, err
	}

	// Apply updates only if fields are provided
	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Category != nil {
		product.Category = *input.Category
	}
	if input.Price != nil {
		product.Price = *input.Price
	}
	if input.Description != nil {
		product.Description = *input.Description
	}
	if input.ImageUrl != nil {
		product.ImageUrl = *input.ImageUrl
	}

	if err := s.repo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}
