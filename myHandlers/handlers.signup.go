package myHandlers

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/database"
	"Golang_Exercises/Golang_Exercises/discovery_project/middlewares"
	"Golang_Exercises/Golang_Exercises/discovery_project/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	db := database.GetUserDB()
	var user models.User
	// var users []models.User
	// var count int = 0
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid credentials")
	}
	// db.Find(&users)
	// for range users {
	// 	count++
	// }
	// user.UserId = count + 1
	if err == nil {
		if middlewares.IsEmail(user.Email) {
			if middlewares.IsPassword(user.Password) {
				err := db.Create(&user).Error
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode("check your data")
					return
				}
				fmt.Println(user)
				json.NewEncoder(w).Encode("Signed Up successfully")
				json.NewEncoder(w).Encode(&user)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("check password should have min 8 characters, should contain  alphabets, numbers and special characters")
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Wrong email")
		}
	}
}
