package myHandlers

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/authentication"
	"encoding/json"
	"net/http"
)

func ValidateJwt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	tokenStr := r.Header.Get("token")

	if authentication.ValidateJwtToken(tokenStr) {
		json.NewEncoder(w).Encode("valid Token")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid Token")
	}
}
