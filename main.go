package main

import (
	"log"
	"net/http"
	"time"

	"backend/src/user"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()

	var userEndpoints = user.MakeEndpoints()

	router.HandleFunc("/users", userEndpoints.Create).Methods("POST")
	router.HandleFunc("/users", userEndpoints.GetAll).Methods("GET")
	// router.HandleFunc("/users/{id}", userEndpoints.Get).Methods("GET")
	// router.HandleFunc("/users/{id}", userEndpoints.Update).Methods("PUT")
	// router.HandleFunc("/users/{id}", userEndpoints.Delete).Methods("DELETE")

	var server = &http.Server{
		Addr:         "127.0.0.1:8000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	var err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
		return
	}
}
