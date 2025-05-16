package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Model
type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

// store
var products []Product = []Product{
	{Id: 101, Name: "Pen", Cost: 10},
	{Id: 102, Name: "Pencil", Cost: 5},
	{Id: 103, Name: "Marker", Cost: 50},
}

type AppRouter struct {
}

// http.Handler interface implementation
func (ar *AppRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello, World!")
	case "/products":
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "error serializing products", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err == nil {
				products = append(products, newProduct)
				w.WriteHeader(http.StatusCreated)
				return
			}
			http.Error(w, "error deserializing payload", http.StatusBadRequest)
		}

	case "/customers":
		fmt.Fprintln(w, "The list of customers will be served!")
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}

}

func main() {
	router := &AppRouter{}
	http.ListenAndServe(":8080", router)
}
