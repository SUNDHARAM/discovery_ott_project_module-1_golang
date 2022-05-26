package myHandlers

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/database"
	"Golang_Exercises/Golang_Exercises/discovery_project/middlewares"
	"Golang_Exercises/Golang_Exercises/discovery_project/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	db := database.GetUserDB()
	var user []models.User
	db.Find(&user)
	fmt.Println(user)
	json.NewEncoder(w).Encode(&user)
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	db := database.GetUserDB()
	err := db.First(&user, "user_id = ?", params["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("user data not exists")
		return
	}
	fmt.Println(user)
	json.NewEncoder(w).Encode(&user)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	db := database.GetUserDB()
	var user models.User
	var users []models.User
	var count int = 0
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid credentials")
	}
	db.Find(&users)
	for range users {
		count++
	}
	user.UserId = count + 1
	if middlewares.IsEmail(user.Email) {
		if middlewares.IsPassword(user.Password) {
			err := db.Create(&user).Error
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("user data already exists")
				return
			}
			fmt.Println(user)
			json.NewEncoder(w).Encode("user added successfully")
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.GetUserDB()
	var user models.User
	params := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid credentials")
	}
	if middlewares.IsEmail(user.Email) {
		if middlewares.IsPassword(user.Password) {
			sqlStatement := `UPDATE public."user" SET email=$1, password=$2, gender=$3, time_stamp=$4, 
				is_active=$5, is_admin=$6, is_cms=$7, is_sms=$8 WHERE user_id=$9`
			db.Exec(sqlStatement, user.Email, user.Password, user.Gender, user.TimeStamp, user.IsActive,
				user.IsAdmin, user.IsCms, user.IsSms, params["id"])
			json.NewEncoder(w).Encode("user details updated successfully")
			json.NewEncoder(w).Encode(&user)
			fmt.Println(user)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("check password should have min 8 characters, should contain  alphabets, numbers and special characters")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Wrong email")
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	var users []models.User
	var del bool = false
	db := database.GetUserDB()

	db.Find(&users)
	for _, index := range users {
		if strconv.Itoa(index.UserId) == params["id"] {
			del = true
			db.Delete(&user, params["id"])
			fmt.Println(user)
			json.NewEncoder(w).Encode("user deleted successfully")
		}
	}
	if !del {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("user not exists")
	}
}
