package authentication

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/models"

	"github.com/dgrijalva/jwt-go"
)

func ValidateJwtToken(tokenstr string) bool {

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenstr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		return false
	}

	return tkn.Valid
}
