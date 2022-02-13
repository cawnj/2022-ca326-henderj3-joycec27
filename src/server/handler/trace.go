package handler

import (
	"net/http"

	"sonic-server/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func trace(router chi.Router) {
	router.Post("/", contactTrace)
}

func contactTrace(w http.ResponseWriter, r *http.Request) {
	userReq := &models.UserRequest{}
	if err := render.Bind(r, userReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	_, err := dbInstance.UpdateCovidPositive(userReq.UserID, true)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	users, err := dbInstance.GetContactUsers(userReq.UserID)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, users); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
