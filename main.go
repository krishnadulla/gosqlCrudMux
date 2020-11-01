package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishnadulla/gosqlCrudMux/handlerfunctions"
)

func main() {
	fmt.Println("Go MySql Tutorial with Mux")

	r := mux.NewRouter()

	r.HandleFunc("/users/create", handlerfunctions.CreateUser).Methods("GET")
	r.HandleFunc("/user/update", handlerfunctions.UpdateUserHandler).Methods("GET")
	r.HandleFunc("/user/delete", handlerfunctions.DeleteUserHandler).Methods("GET")
	r.HandleFunc("/users/all", handlerfunctions.FetchAllUsers).Methods("GET")

	r.HandleFunc("/user/{id}", handlerfunctions.FetchUserByParams).Methods("GET")

	http.ListenAndServe(":8080", r)

}
