package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPass       string `mapstructure:"DB_PASS"`
	DBName       string `mapstructure:"DB_NAME"`
	ClientId     string `mapstructure:"COGNITO_APP_CLIENT_ID"`
	ClientSecret string `mapstructure:"COGNITO_APP_CLIENT_SECRET"`
	UserPollId   string `mapstructure:"COGNITO_USER_POOL_ID"`
	Region       string `mapstructuree:"AWS_REGION"`
}

func LoadEnv() Env {
	env := Env{}

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env: ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return env
}
