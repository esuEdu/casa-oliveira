package service

import (
	"github.com/esuEdu/casa-oliveira/internal/dto"
)

type AuthService interface {
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) SignIn(c dto.SignupRequest) error {

	return nil
}

func (s *authService) SingUp(c dto.SignupRequest) error {

	return nil
}
