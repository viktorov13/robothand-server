package jwtutil

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("SUPER_SECRET_KEY")

func GenerateToken(uuid string) (string, error) {
	claims := jwt.MapClaims{
		"uuid": uuid,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}
