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

/* *********************************************************** */
// Generic Router implementation

type AppRouter struct {
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func NewAppRouter() *AppRouter {
	return &AppRouter{
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

func (ar *AppRouter) Register(resourcePath string, handlerFn func(http.ResponseWriter, *http.Request)) {
	ar.routes[resourcePath] = handlerFn
}

// http.Handler interface implementation
func (ar *AppRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	if handlerFn, exists := ar.routes[r.URL.Path]; exists {
		handlerFn(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

/* *********************************************************** */
// Application specific routes
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
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
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The list of customers will be served!")
}

func main() {
	router := NewAppRouter()
	router.Register("/", IndexHandler)
	router.Register("/products", ProductsHandler)
	router.Register("/customers", CustomersHandler)
	http.ListenAndServe(":8080", router)
}
