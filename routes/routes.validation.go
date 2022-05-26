package routes

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/myHandlers"

	"github.com/gorilla/mux"
)

func RouteValidation(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/validation", myHandlers.ValidateJwt).Methods("POST")
}
