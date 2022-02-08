package handler

import (
	"net/http"

	"sonic-server/db"

	"github.com/go-chi/chi"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.Route("/users", users)
	router.Route("/locations", locations)
	router.Route("/entrylog", entrylog)
	router.Route("/trace", trace)
	return router
}
