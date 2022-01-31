package models

import (
	"fmt"
	"net/http"

	geo "github.com/kellydunn/golang-geo"
)

type Location struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Coords geo.Point `json:"coords"`
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
