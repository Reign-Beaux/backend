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
)

func MakeEndpoints() Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(),
		Get:    makeGetEndpoint(),
		GetAll: makeGetAllEndpoint(),
		Update: makeUpdateEndpoint(),
		Delete: makeDeleteEndpoint(),
	}
}

func makeCreateEndpoint() Controller {
	return func(writer http.ResponseWriter, request *http.Request) {

		var createRequest CreateRequest

		var error = json.NewDecoder(request.Body).Decode(&createRequest)

		if error != nil {
			http.Error(writer, error.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	}
}

func makeGetEndpoint() Controller {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("get user")
		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	}
}

func makeGetAllEndpoint() Controller {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("get all users")
		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	}
}

func makeUpdateEndpoint() Controller {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("update user")
		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	}
}

func makeDeleteEndpoint() Controller {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("delete user")
		json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	}
}
