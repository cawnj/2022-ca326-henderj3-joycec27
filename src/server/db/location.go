package db

import (
	"database/sql"

	"sonic-server/models"
)

func (db Database) GetLocation(locationId int) (*models.Location, error) {
	location := &models.Location{}
	query := `SELECT * FROM locations
	WHERE location_id = $1`
	row := db.Conn.QueryRow(query, locationId)
	err := row.Scan(&location.LocationID, &location.Name, &location.Coords)
	switch err {
	case sql.ErrNoRows:
		return nil, ErrNoMatch
	case nil:
		return location, nil
	default:
		return nil, err
	}
}

func (db Database) GetAllLocations() (*models.LocationList, error) {
	locations := &models.LocationList{}
	rows, err := db.Conn.Query("SELECT * FROM locations ORDER BY location_id DESC")
	if err != nil {
		return locations, err
	}
	for rows.Next() {
		var location models.Location
		err := rows.Scan(&location.LocationID, &location.Name, &location.Coords)
		if err != nil {
			return locations, err
		}
		locations.Locations = append(locations.Locations, location)
	}
	return locations, nil
}
