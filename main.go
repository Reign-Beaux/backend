package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"backend/src/user"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var router = mux.NewRouter()
	_ = godotenv.Load()
	var logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	var dsn = fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)
	var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()

	var _ = db.AutoMigrate(&user.User{})

	var userRepository = user.NewRepository(logger, db)
	var userService = user.NewService(logger, userRepository)
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
