package service

import (
	"github.com/esuEdu/casa-oliveira/internal/entity"
	"github.com/esuEdu/casa-oliveira/internal/repositories"
)

type UserService interface {
}

type userService struct {
	repo repositories.UserRepo
}

func NewUserService(r repositories.UserRepo) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) Create(u *entity.User, password string) error {

	return nil
}
