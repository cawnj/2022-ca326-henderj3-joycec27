package models

import geo "github.com/kellydunn/golang-geo"

type Location struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Coords geo.Point `json:"coords"`
}
