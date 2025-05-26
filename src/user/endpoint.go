package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(writer http.ResponseWriter, request *http.Request)
	Endpoints  struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}
	CreateRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	ErrorResponse struct {
		Error string `json:"error"`
	}
)

func MakeEndpoints(service Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(service),
		Get:    makeGetEndpoint(service),
		GetAll: makeGetAllEndpoint(service),
		Update: makeUpdateEndpoint(service),
		Delete: makeDeleteEndpoint(service),
	}
}

func makeCreateEndpoint(service Service) Controller {
	return func(writer http.ResponseWriter, request *http.Request) {

		var createRequest CreateRequest

		var err = json.NewDecoder(request.Body).Decode(&createRequest)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			json.NewEncoder(writer).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		if createRequest.FirstName == "" {
			http.Error(writer, "first name is required", http.StatusBadRequest)
			json.NewEncoder(writer).Encode(ErrorResponse{Error: "first name is required"})
			return
		}

		if createRequest.LastName == "" {
			http.Error(writer, "last name is required", http.StatusBadRequest)
			json.NewEncoder(writer).Encode(ErrorResponse{Error: "last name is required"})
			return
		}

		var serviceErr = service.Create(createRequest.FirstName, createRequest.LastName, createRequest.Email, createRequest.Phone)

		if serviceErr != nil {
			http.Error(writer, serviceErr.Error(), http.StatusBadRequest)
			json.NewEncoder(writer).Encode(ErrorResponse{Error: serviceErr.Error()})
			return
		}

		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	}
}

func makeGetEndpoint(service Service) Controller {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("get user")
		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	}
}

func makeGetAllEndpoint(service Service) Controller {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("get all users")
		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	}
}

func makeUpdateEndpoint(service Service) Controller {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("update user")
		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	}
}

func makeDeleteEndpoint(service Service) Controller {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("delete user")
		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	}
}
