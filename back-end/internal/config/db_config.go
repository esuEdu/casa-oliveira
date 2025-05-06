package config

import (
	"fmt"

	"github.com/esuEdu/casa-oliveira/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	env := LoadEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.DBHost,
		env.DBUser,
		env.DBPass,
		env.DBName,
		env.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	return db
}
