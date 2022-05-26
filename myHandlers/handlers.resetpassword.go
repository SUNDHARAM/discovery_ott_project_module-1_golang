package myHandlers

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/database"
	"Golang_Exercises/Golang_Exercises/discovery_project/middlewares"
	"Golang_Exercises/Golang_Exercises/discovery_project/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var myresetpwd models.ResetPwd
	err := json.NewDecoder(r.Body).Decode(&myresetpwd)
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
		if index.Email == myresetpwd.Email {
			checkemail = true
			if index.Password == myresetpwd.OldPassword {
				check = true
				if middlewares.IsPassword(myresetpwd.NewPassword) {
					if myresetpwd.NewPassword != myresetpwd.OldPassword {
						sqlStatement := `UPDATE public."user" SET password=$1 WHERE user_id=$2`
						db.Exec(sqlStatement, myresetpwd.NewPassword, index.UserId)

						index.Password = myresetpwd.NewPassword
						fmt.Println(index)
						json.NewEncoder(w).Encode("New password updated")
						json.NewEncoder(w).Encode(&index)
					} else {
						w.WriteHeader(http.StatusUnauthorized)
						json.NewEncoder(w).Encode("New password and old password are same")
					}
				} else {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode("check password should have min 8 characters, should contain  alphabets, numbers and special characters")
				}
			}
		}
	}
	if !checkemail && err == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("wrong email")
	}

	if !check && checkemail {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("wrong old password")
	}

}
