package handler

import (
	"fmt"
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
	users, err := dbInstance.GetContactUsers(userReq.UserID)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := notifyCloseContacts(users); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
	response := &models.PostResponse{
		StatusCode: 201,
		StatusText: fmt.Sprintf("notified '%d' close contacts", len(users.Users)),
	}
	if err := render.Render(w, r, response); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func notifyCloseContacts(users *models.UserList) error {
	return nil
}
