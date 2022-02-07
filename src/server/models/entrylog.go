package models

import (
	"database/sql"
	"fmt"
	"net/http"
)

type EntryLog struct {
	EntryID    int    `json:"entry_id"`
	UserID     int    `json:"user_id"`
	LocationID int    `json:"location_id"`
	EntryTime  string `json:"entry_time"`
	ExitTime   string `json:"exit_time"`
}

type DBEntryLog struct {
	EntryID    int            `json:"entry_id"`
	UserID     int            `json:"user_id"`
	LocationID int            `json:"location_id"`
	EntryTime  sql.NullString `json:"entry_time"`
	ExitTime   sql.NullString `json:"exit_time"`
}

type EntryLogList struct {
	EntryLogs []EntryLog `json:"entry_logs"`
}

func (u *EntryLog) Bind(r *http.Request) error {
	if u.UserID == 0 {
		return fmt.Errorf("user_id is a required field")
	}
	if u.LocationID == 0 {
		return fmt.Errorf("location_id is a required field")
	}
	return nil
}

func (*EntryLogList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*EntryLog) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*DBEntryLog) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
