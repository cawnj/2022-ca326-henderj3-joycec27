package models

import (
	"fmt"
	"net/http"
)

type EntryLogRequest struct {
	UserID     string `json:"user_id"`
	LocationID int    `json:"location_id"`
	Timestamp  string `json:"timestamp"`
}

func (u *EntryLogRequest) Bind(r *http.Request) error {
	if u.UserID == "" {
		return fmt.Errorf("user_id is a required field")
	}
	if u.LocationID == 0 {
		return fmt.Errorf("location_id is a required field")
	}
	if u.Timestamp == "" {
		return fmt.Errorf("timestamp is a required field")
	}
	return nil
}

func (*EntryLogRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
