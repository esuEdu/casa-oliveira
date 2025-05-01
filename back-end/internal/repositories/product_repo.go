package repositories

import (
	"strconv"

	"github.com/esuEdu/casa-oliveira/internal/dto"
	"github.com/esuEdu/casa-oliveira/internal/entity"
	"gorm.io/gorm"
)

type ProductRepo interface {
	Create(p *entity.Product) error
	FindByID(id uint) (*entity.Product, error)
	List(page, pageSize int) (*dto.ProductPagination, error)
	Update(p *entity.Product) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) Create(p *entity.Product) error {

	// Validate that price is a valid decimal string
	if _, err := strconv.ParseFloat(p.Price, 64); err != nil {
		return err
	}

	if err := r.db.Create(p).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepo) List(page, pageSize int) (*dto.ProductPagination, error) {

	var products []entity.Product
	var total int64

	query := r.db.Model(&entity.Product{})

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	offset := (page - 1) * pageSize
	if err := query.
		Limit(pageSize).
		Offset(offset).
		Find(&products).
		Error; err != nil {
		return nil, err
	}

	return &dto.ProductPagination{
		Results:  products,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func (r *productRepo) FindByID(id uint) (*entity.Product, error) {
	var product entity.Product

	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) Update(p *entity.Product) error {
	return r.db.Save(p).Error
}
