package main

import (
	"fmt"
	"log"
	"net/http"

	"sonic-server/db"
	"sonic-server/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Alive!")
	})
	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		data := &models.User{}
		if err := render.Bind(r, data); err != nil {
			fmt.Fprint(w, "Error")
			return
		}
		render.JSON(w, r, data)
	})

	log.Println("Attemping connection to database")
	_, err := db.Initialize()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	log.Println("Server starting at http://localhost:8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
