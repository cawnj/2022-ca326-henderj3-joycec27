package models

import (
	"fmt"
	"net/http"
)

type Location struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Coords string `json:"coords"`
}

type LocationList struct {
	Locations []Location `json:"locations"`
}

func (u *Location) Bind(r *http.Request) error {
	if u.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (*LocationList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Location) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
