package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omit"`
	Phone    string `json:"phone"`
	Role     string `json:"role" gorm:"default:'user'"`
}
