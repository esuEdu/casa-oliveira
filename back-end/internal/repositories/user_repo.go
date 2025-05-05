package repositories

import (
	"github.com/esuEdu/casa-oliveira/internal/entity"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(u *entity.User) error
	FindByID(id uint) (*entity.User, error)
	Update(u *entity.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(u *entity.User) error {
	if err := r.db.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) FindByID(id uint) (*entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Update(u *entity.User) error {
	return r.db.Save(u).Error
}
