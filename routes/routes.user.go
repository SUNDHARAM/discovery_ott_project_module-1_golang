package routes

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/myHandlers"

	"github.com/gorilla/mux"
)

func RouteUser(mainRouter *mux.Router) {
	userRouter := mainRouter.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/getallusers", myHandlers.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/getoneuser/{id}", myHandlers.GetOneUser).Methods("GET")
	userRouter.HandleFunc("/adduser", myHandlers.AddUser).Methods("POST")
	userRouter.HandleFunc("/updateuser/{id}", myHandlers.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/deleteuser/{id}", myHandlers.DeleteUser).Methods("DELETE")

}
