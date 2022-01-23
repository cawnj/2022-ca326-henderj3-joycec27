package main

import (
	"fmt"
	"log"
	"net/http"

	"sonic-server/db"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Alive!")
	})

	log.Println("Attemping connection to database")
	_, err := db.Initialize()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	log.Println("Server starting at http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
