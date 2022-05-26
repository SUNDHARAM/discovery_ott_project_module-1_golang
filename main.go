package main

import (
	"Golang_Exercises/Golang_Exercises/discovery_project/database"
	"Golang_Exercises/Golang_Exercises/discovery_project/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("discovery project started")
	database.Connectmydatabase()
	defer database.Closemydatabase()

	mainRouter := mux.NewRouter()

	routes.RouteUser(mainRouter)
	routes.RouteSignUp(mainRouter)
	routes.RouteLogin(mainRouter)
	routes.RouteResetPassword(mainRouter)
	routes.RouteValidation(mainRouter)

	fmt.Printf("server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", mainRouter))

}
