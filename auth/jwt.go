package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":userID,
	})

	tokenString, err := token.SignedString([]byte("secretkey"))
	return tokenString, err
}
