package models

import (
	"fmt"
	"net/http"
)

type Location struct {
	LocationID int    `json:"location_id"`
	Name       string `json:"name"`
	Coords     string `json:"coords"`
}

type LocationList struct {
	Locations []Location `json:"locations"`
}

type LocationNameAndTimestamp struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
}

func (u *Location) Bind(r *http.Request) error {
	if u.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (*Location) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*LocationList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*LocationNameAndTimestamp) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
