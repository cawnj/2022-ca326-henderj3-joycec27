package models

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

type EntryLog struct {
	EntryID    int    `json:"entry_id"`
	UserID     string `json:"user_id"`
	LocationID int    `json:"location_id"`
	EntryTime  string `json:"entry_time"`
	ExitTime   string `json:"exit_time"`
}

type EntryLogList struct {
	EntryLogs []EntryLog `json:"entry_logs"`
}

func (u *EntryLog) Bind(r *http.Request) error {
	if u.EntryID == 0 {
		return fmt.Errorf("entry_id is a required field")
	}
	if u.UserID == "" {
		return fmt.Errorf("user_id is a required field")
	}
	if u.LocationID == 0 {
		return fmt.Errorf("location_id is a required field")
	}
	return nil
}

func (e *EntryLog) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)
	return nil
}

func (*EntryLogList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
