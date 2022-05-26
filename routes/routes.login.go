package routes

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/myHandlers"

	"github.com/gorilla/mux"
)

func RouteLogin(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/login", myHandlers.Login).Methods("POST")
}
