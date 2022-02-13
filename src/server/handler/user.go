package handler

import (
	"net/http"

	"sonic-server/db"
	"sonic-server/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func user(router chi.Router) {
	router.Get("/", getUser)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userReq := &models.UserRequest{}
	if err := render.Bind(r, userReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	user, err := dbInstance.GetUser(userReq.UserID)
	switch {
	case err == db.ErrNoMatch:
		render.Render(w, r, ErrNotFound)
		return
	case err != nil:
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
