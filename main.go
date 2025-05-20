package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/courses", getCourses).Methods("GET")

	var server = &http.Server{
		Addr:         "127.0.0.1:8000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	var error = server.ListenAndServe()

	if error != nil {
		log.Fatal(error)
		return
	}
}

func getUsers(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
}

func getCourses(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
}
