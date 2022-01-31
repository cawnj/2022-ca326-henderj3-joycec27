package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func locations(router chi.Router) {
	router.Get("/", getAllLocations)
}

func getAllLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := dbInstance.GetAllLocations()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
	}
	if err := render.Render(w, r, locations); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
