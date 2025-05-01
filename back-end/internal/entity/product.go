package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string   `json:"name"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Price       string   `jsonn:"price" gorm:"type:decimal(19,4)"`
	ImageUrl    []string `json:"image_url,omitempty" gorm:"type:jsonb;serializer:json"`
}
