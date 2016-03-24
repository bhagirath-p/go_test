package controllers

import (
	"log"
	"net/http"
)

type RegisterController struct {}

var Register RegisterController

func (r RegisterController) Create(rw http.ResponseWriter, req *http.Request) {
	log.Println("here")
}

// func Create(req Request, rw ResponseWriter) {

// }