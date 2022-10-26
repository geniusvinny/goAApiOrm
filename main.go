package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// match our methods against particular paths and http variables
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", Deleteuser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("GO TO THE WORLD, --- main.go --- is running ")
	// call the initial migration whicg=h will create our tables for us
	// the table is created only if it does not already exist in the database
	InitialMigration()

	handleRequests()

}
