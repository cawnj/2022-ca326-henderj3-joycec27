package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Test struct {
	Message string `json:"message"`
}

func trace(router chi.Router) {
	router.Post("/", contactTrace)
}

func contactTrace(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &Test{
		Message: "hello",
	})
}

func (t *Test) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, 200)
	return nil
}
