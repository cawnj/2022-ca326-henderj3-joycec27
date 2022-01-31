package db

import (
	"sonic-server/models"
)

func (db Database) GetAllLocations() (*models.LocationList, error) {
	locations := &models.LocationList{}
	rows, err := db.Conn.Query("SELECT * FROM locations ORDER BY location_id DESC")
	if err != nil {
		return locations, err
	}
	for rows.Next() {
		var location models.Location
		err := rows.Scan(&location.ID, &location.Name, &location.Coords)
		if err != nil {
			return locations, err
		}
		locations.Locations = append(locations.Locations, location)
	}
	return locations, nil
}
