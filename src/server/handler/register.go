package handler

import (
	"fmt"
	"net/http"

	"sonic-server/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func register(router chi.Router) {
	router.Post("/", registerUser)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.RegisterUser(user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	response := &models.PostResponse{
		StatusCode: 201,
		StatusText: fmt.Sprintf("user '%s' registered successfuly", user.UserID),
	}
	if err := render.Render(w, r, response); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
