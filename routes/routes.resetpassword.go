package routes

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/myHandlers"

	"github.com/gorilla/mux"
)

func RouteResetPassword(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/resetpwd", myHandlers.ResetPassword).Methods("POST")
}
