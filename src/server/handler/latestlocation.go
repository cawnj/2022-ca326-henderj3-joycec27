package handler

import (
	"net/http"

	"sonic-server/db"
	"sonic-server/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func latestLocation(router chi.Router) {
	router.Get("/", getLatestLocation)
}

func getLatestLocation(w http.ResponseWriter, r *http.Request) {
	userReq := &models.UserRequest{}
	if err := render.Bind(r, userReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	// ensure user exists
	user, err := dbInstance.GetUser(userReq.UserID)
	switch {
	case err == db.ErrNoMatch:
		render.Render(w, r, ErrNotFound)
		return
	case err != nil:
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	// get their latest entry log
	latestEntryLog, err := dbInstance.GetLatestEntryLog(user.UserID)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	// get the location data
	latestLocation, err := dbInstance.GetLocation(latestEntryLog.LocationID)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	// build response object
	locationNameAndTimestamp := &models.LocationNameAndTimestamp{
		Name:      latestLocation.Name,
		Timestamp: latestEntryLog.EntryTime,
	}
	if err := render.Render(w, r, locationNameAndTimestamp); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
