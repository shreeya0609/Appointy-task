package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/user", CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))

}
