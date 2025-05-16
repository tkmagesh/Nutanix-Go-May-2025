package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
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
	if handlerFn, exists := ar.routes[r.URL.Path]; exists {
		handlerFn(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

/* *********************************************************** */
// Application specific routes
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[IndexHandler] trace id : ", r.Context().Value("trace-id"))
	time.Sleep(6 * time.Second)
	if r.Context().Err() == nil {
		fmt.Fprintln(w, "Hello, World!")
		return
	}
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

// middlewares
func logMiddleware(handlerFn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("trace-id : [%v], method : [%s] - url : %s\n", r.Context().Value("trace-id"), r.Method, r.URL.Path)
		handlerFn(w, r)
	}
}

func traceMiddleware(handlerFn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		traceId := rand.Intn(10000)
		traceCtx := context.WithValue(r.Context(), "trace-id", traceId)
		handlerFn(w, r.WithContext(traceCtx))
	}
}

func timeoutMiddleware(handlerFn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		timeoutCtx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()
		handlerFn(w, r.WithContext(timeoutCtx))
		if timeoutCtx.Err() == context.DeadlineExceeded {
			fmt.Fprintln(w, "request timedout", http.StatusRequestTimeout)
		}
	}
}

func main() {
	router := NewAppRouter()
	logIndexHandler := timeoutMiddleware(traceMiddleware(logMiddleware(IndexHandler)))
	router.Register("/", logIndexHandler)

	router.Register("/products", timeoutMiddleware(traceMiddleware(logMiddleware(ProductsHandler))))
	router.Register("/customers", timeoutMiddleware(traceMiddleware(logMiddleware(CustomersHandler))))
	http.ListenAndServe(":8080", router)
}
