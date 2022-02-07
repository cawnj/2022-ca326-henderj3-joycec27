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
	entryLog := &models.EntryLog{} // TODO: change to request struct
	if err := render.Bind(r, entryLog); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	latestEntryLog, err := dbInstance.GetLatestEntryLog(entryLog.UserID)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if latestEntryLog.ExitTime.Valid { // is the exit time not NULL
		createEntryLog(w, r, entryLog)
	} else {
		updateEntryLog(w, r, entryLog)
	}

	if err := render.Render(w, r, entryLog); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func createEntryLog(w http.ResponseWriter, r *http.Request, entryLog *models.EntryLog) {
	if err := dbInstance.AddEntryLog(entryLog); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
}

func updateEntryLog(w http.ResponseWriter, r *http.Request, entryLog *models.EntryLog) {
	if err := dbInstance.UpdateEntryLog(entryLog); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
}
