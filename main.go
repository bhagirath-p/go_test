package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/bhagirath-p/go_test/api/v1/controllers"
)

func main() {
	r:=mux.NewRouter()
	log.Println("here main")
	r.HandleFunc("/sign_up", controllers.Register.Create).Methods("POST")
	http.Handle("/", r)
	log.Println("main : Started : Listening on : http://localhost:3000 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}