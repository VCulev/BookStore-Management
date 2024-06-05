package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VCulev/BookStore-Management/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Running server on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
