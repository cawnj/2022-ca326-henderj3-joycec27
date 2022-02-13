package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func users(router chi.Router) {
	router.Get("/", getAllUsers)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dbInstance.GetAllUsers()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
	}
	if err := render.Render(w, r, users); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
