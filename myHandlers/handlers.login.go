package myHandlers

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/authentication"
	"Golang_Exercises/Golang_Exercises/discovery_project/database"
	"Golang_Exercises/Golang_Exercises/discovery_project/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var credentials models.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid credentials")
	}
	db := database.GetUserDB()
	var user []models.User
	db.Find(&user)

	var check bool = false
	var checkemail bool = false

	for _, index := range user {
		if index.Email == credentials.Email {
			checkemail = true
			if index.Password == credentials.Password {
				check = true
				tokenString := authentication.GenerateJwtToken(credentials.Email)
				fmt.Println(index)
				fmt.Println(tokenString)
				json.NewEncoder(w).Encode(&tokenString)
			}
		}
	}
	if !checkemail && err == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("wrong email")
	}

	if !check && checkemail {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("wrong password")
	}
}
