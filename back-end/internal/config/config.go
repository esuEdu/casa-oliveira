package config

import (
	"log"
	"os"
	"strconv"
)

type Env struct {
	AppPort                string
	DBHost                 string
	DBPort                 string
	DBUser                 string
	DBPass                 string
	DBName                 string
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
	AccessTokenSecret      string
	RefreshTokenSecret     string
}

func LoadEnv() Env {
	return Env{
		AppPort:                getEnvString("APP_PORT", false),
		DBHost:                 getEnvString("DB_HOST", true),
		DBPort:                 getEnvString("DB_PORT", true),
		DBUser:                 getEnvString("DB_USER", true),
		DBPass:                 getEnvString("DB_PASSWORD", true),
		DBName:                 getEnvString("DB_NAME", true),
		AccessTokenExpiryHour:  getEnvInt("ACCESS_TOKEN_EXPIRY_HOUR", true),
		RefreshTokenExpiryHour: getEnvInt("REFRESH_TOKEN_EXPIRY_HOUR", true),
		AccessTokenSecret:      getEnvString("ACCESS_TOKEN_SECRET", true),
		RefreshTokenSecret:     getEnvString("REFRESH_TOKEN_SECRET", true),
	}
}

// Helper functions
func getEnvString(key string, required bool) string {
	value := os.Getenv(key)
	if required && value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}

func getEnvInt(key string, required bool) int {
	strValue := getEnvString(key, required)
	if strValue == "" && !required {
		return 0
	}

	value, err := strconv.Atoi(strValue)
	if err != nil {
		log.Fatalf("Invalid integer value for %s: %v", key, err)
	}
	return value
}
