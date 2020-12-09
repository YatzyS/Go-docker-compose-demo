package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	postRouter := r.Methods("POST").Subrouter()
	postRouter.HandleFunc("/scrap", Scrap)
	postRouter.HandleFunc("/add", Add)
}

//Scrap data from amazon
func Scrap(rw http.ResponseWriter, r *http.Request) {

}

//Add data from amazon to DB
func Add(rw http.ResponseWriter, r *http.Request) {

}