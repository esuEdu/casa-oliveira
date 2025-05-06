package service

import (
	"errors"
	"strconv"

	"github.com/esuEdu/casa-oliveira/internal/dto"
	"github.com/esuEdu/casa-oliveira/internal/entity"
	"github.com/esuEdu/casa-oliveira/internal/repositories"
	"github.com/esuEdu/casa-oliveira/internal/util"
)

type UserService interface {
	SignUp(u *dto.UserDTO, password string) (*dto.UserDTO, error)
	SignIn(email, password string) (dto.AuthDTO, error)
}

type userService struct {
	repo repositories.UserRepo
}

func NewUserService(r repositories.UserRepo) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) SignUp(u *dto.UserDTO, password string) (*dto.UserDTO, error) {
	existingUser, _ := s.repo.FindByEmail(u.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	stringHash, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     u.Name,
		Email:    u.Email,
		Phone:    u.Phone,
		Password: stringHash,
	}

	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}, nil

}

func (s *userService) SignIn(email, password string) (dto.AuthDTO, error) {

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return dto.AuthDTO{}, errors.New("user not found")
	}

	err = util.CheckPasswordHash(password, user.Password)
	if err != nil {
		return dto.AuthDTO{}, errors.New("wrong credential")
	}

	token, err := util.GenerateToken(strconv.FormatUint(uint64(user.ID), 10), user.Role)
	if err != nil {
		return dto.AuthDTO{}, errors.New("failed creating token")
	}

	return dto.AuthDTO{
		Token: token,
		Role:  user.Role,
	}, nil
}
