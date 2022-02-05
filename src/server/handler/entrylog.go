package handler

import (
	"net/http"

	"sonic-server/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func entrylog(router chi.Router) {
	router.Post("/", writeEntryLog)
}

func writeEntryLog(w http.ResponseWriter, r *http.Request) {
	entryLog := &models.EntryLog{}
	if err := render.Bind(r, entryLog); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddEntryLog(entryLog); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, entryLog); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
