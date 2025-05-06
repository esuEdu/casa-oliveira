package dto

type AuthDTO struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}
