package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/VCulev/BookStore-Management/pkg/config"
	"github.com/VCulev/BookStore-Management/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	if err := waitForDB(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Running server on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}

func waitForDB() error {
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		_, err := config.Connect()
		if err != nil {
			fmt.Printf("Failed to connect to the database. Retrying in 2 seconds... (attempt %d/%d)\n", i+1, maxRetries)
			time.Sleep(2 * time.Second)
			continue
		}
		fmt.Println("Database is running or started.")
		return nil
	}
	return fmt.Errorf("could not connect to the database after %d attempts", maxRetries)
}
