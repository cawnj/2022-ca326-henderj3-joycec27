package handler

import (
	"fmt"
	"net/http"

	"sonic-server/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func register(router chi.Router) {
	router.Get("/", registerUser)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	registerReq := &models.RegisterRequest{}
	if err := render.Bind(r, registerReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.RegisterUser(registerReq); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	response := &models.RegisterResponse{
		StatusCode: 201,
		StatusText: fmt.Sprintf("user '%s' registered successfuly", registerReq.UserID),
	}
	if err := render.Render(w, r, response); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
