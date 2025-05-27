package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"backend/src/user"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var router = mux.NewRouter()

	var dsn = fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"root",
		"127.0.0.1",
		"3320",
		"go-course-web",
	)
	var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()

	var _ = db.AutoMigrate(&user.User{})

	var userService = user.NewService()
	var userEndpoints = user.MakeEndpoints(userService)

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
