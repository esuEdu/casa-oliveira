package util

import (
	"errors"
	"time"

	"github.com/esuEdu/casa-oliveira/internal/config"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID, role string) (string, error) {

	env := config.LoadEnv()

	mySigningKey := []byte(env.AccessTokenSecret)

	claims := &jwt.MapClaims{
		"user_id":   userID,
		"user_role": role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), //token valid for 24h
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {

	env := config.LoadEnv()

	mySigningKey := []byte(env.AccessTokenSecret)

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing invalid")
		}
		return mySigningKey, nil
	})

	return token, err

}
