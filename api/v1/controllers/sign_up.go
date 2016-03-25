package controllers

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/bhagirath-p/go_test/api/v1/models"
	"io/ioutil"
)

type RegisterController struct {}

var Register RegisterController

func (r RegisterController) Create(rw http.ResponseWriter, req *http.Request) {
	var u models.User
	req_body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	body := json.Unmarshal(req_body, &u)
	if body != nil {
		panic(err)
	}
	log.Printf("%v", u)
}

// func Create(req Request, rw ResponseWriter) {

// }