package authentication

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

func GenerateJwtToken(email string) string {

	expirationTime := time.Now().Add(time.Minute * 10)

	claims := &models.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		panic("problem in generating token")
	}
	return tokenString
}
