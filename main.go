package main

import (
	"log"
	"github.com/gorilla/mux"
	"github.com/bhagirath-p/go_test/api/v1/controllers"
)

func main() {
	r:=mux.NewRouter()
	log.Println("here main")
	r.HandleFunc("/sign_up", controllers.Register.Create).Methods("POST")
}