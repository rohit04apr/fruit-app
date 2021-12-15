package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getFruits(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	json.NewEncoder(response).Encode(fruit)
}

func getFruit(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	for _, item := range fruit {
		if item.Name == params["name"] {
			json.NewEncoder(response).Encode(item)
		}
	}
}

func addFruit(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var newFruit Fruits
	json.NewDecoder(request.Body).Decode(&newFruit)
	newFruit.ID = strconv.Itoa(len(fruit) + 1)
	fruit = append(fruit, newFruit)
	json.NewEncoder(response).Encode(fruit)
}

func updateFruit(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	for i, item := range fruit {
		if item.ID == params["id"] {
			fruit = append(fruit[:i], fruit[i+1:]...)
			var newFruit Fruits
			json.NewDecoder(request.Body).Decode(&newFruit)
			newFruit.ID = params["id"]
			fruit = append(fruit, newFruit)
			json.NewEncoder(response).Encode(fruit)
			return

		}
	}
}

func deleteFruit(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	for i, item := range fruit {
		if item.ID == params["id"] {
			fruit = append(fruit[:i], fruit[i+1:]...)
			break

		}
	}
}

type Fruits struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var fruit []Fruits

func main() {

	fruit = append(fruit, Fruits{ID: "1", Name: "Mangos", Price: 2.99})
	fruit = append(fruit, Fruits{ID: "2", Name: "Apples", Price: 3.99})
	fruit = append(fruit, Fruits{ID: "3", Name: "Grapes", Price: 4.99})

	handler := mux.NewRouter()

	handler.HandleFunc("/fruits", getFruits).Methods("GET")
	handler.HandleFunc("/fruits/{name}", getFruit).Methods("GET")
	handler.HandleFunc("/fruits", addFruit).Methods("POST")
	handler.HandleFunc("/fruits/{id}", updateFruit).Methods("PUT")
	handler.HandleFunc("/fruits/{id}", deleteFruit).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", handler))
}
