package handler

import (
	"fmt"
	"net/http"

	"sonic-server/db"

	"github.com/go-chi/chi"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	dbInstance = db
	router := chi.NewRouter()
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Alive!")
	})

	router.Route("/user", user)
	router.Route("/users", users)
	router.Route("/locations", locations)
	router.Route("/entrylog", entrylog)
	router.Route("/trace", trace)
	router.Route("/register", register)
	router.Route("/latestlocation", latestLocation)
	return router
}
