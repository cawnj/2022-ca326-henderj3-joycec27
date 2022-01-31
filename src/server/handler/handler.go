package handler

import (
	"net/http"

    "github.com/go-chi/chi"
	"sonic-server/db"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.Route("/users", users)
	return router
}
