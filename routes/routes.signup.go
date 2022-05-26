package routes

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/myHandlers"

	"github.com/gorilla/mux"
)

func RouteSignUp(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/signup", myHandlers.SignUp).Methods("POST")
}
