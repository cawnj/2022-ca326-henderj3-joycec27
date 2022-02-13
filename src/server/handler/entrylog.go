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
	entryLogReq := &models.EntryLogRequest{}
	if err := render.Bind(r, entryLogReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	latestEntryLog, err := dbInstance.GetLatestEntryLog(entryLogReq.UserID)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	entryLog := &models.EntryLog{
		UserID:     entryLogReq.UserID,
		LocationID: entryLogReq.LocationID,
	}

	// if this user does not have an entry log OR
	//   this user's latest entry log includes an exit time
	if latestEntryLog == nil || latestEntryLog.ExitTime != "" {
		entryLog.EntryTime = entryLogReq.Timestamp
		createEntryLog(w, r, entryLog)
	} else {
		entryLog.ExitTime = entryLogReq.Timestamp
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
