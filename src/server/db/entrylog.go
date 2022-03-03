package db

import (
	"time"

	"database/sql"

	"sonic-server/models"
)

func (db Database) AddEntryLog(entryLog *models.EntryLog) error {
	var id int
	query := `INSERT INTO entry_log (user_id, location_id, entry_time)
	VALUES ($1, $2, $3::timestamp)
	RETURNING entry_id`
	err := db.Conn.QueryRow(
		query,
		entryLog.UserID,
		entryLog.LocationID,
		entryLog.EntryTime,
	).Scan(&id)
	if err != nil {
		return err
	}
	entryLog.EntryID = id
	return nil
}

func (db Database) UpdateEntryLog(entryLog *models.EntryLog) error {
	var id int
	var entryTime string
	query := `UPDATE entry_log
	SET exit_time = $2::timestamp
	WHERE entry_id = (
		SELECT max(entry_id)
		FROM entry_log
		WHERE user_id = $1
	)
	RETURNING entry_id, entry_time`
	row := db.Conn.QueryRow(query, entryLog.UserID, entryLog.ExitTime)
	err := row.Scan(&id, &entryTime)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	case nil:
		entryLog.EntryID = id
		entryLog.EntryTime = entryTime
		return nil
	default:
		return err
	}
}

func (db Database) GetLatestEntryLog(userId string) (*models.EntryLog, error) {
	var entryId int
	var locationId int
	var entryTime string        // will never be null
	var exitTime sql.NullString // sometimes null
	query := `SELECT entry_id, location_id, entry_time, exit_time FROM entry_log
	WHERE entry_id = (
		SELECT max(entry_id)
		FROM entry_log
		WHERE user_id = $1
	)`
	err := db.Conn.QueryRow(query, userId).Scan(
		&entryId,
		&locationId,
		&entryTime,
		&exitTime,
	)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		entryLog := &models.EntryLog{
			EntryID:    entryId,
			UserID:     userId,
			LocationID: locationId,
			EntryTime:  entryTime,
			ExitTime:   exitTime.String,
		}
		return entryLog, nil
	}
}

func (db Database) GetContactUsers(userId string) (*models.UserList, error) {
	dt := time.Now().AddDate(0, 0, -3).Format("2006-01-02 15:04:05")
	contactUserIds := make(map[string]bool)
	contactUsers := &models.UserList{}
	query := `SELECT location_id, entry_time, exit_time FROM entry_log
	WHERE user_id = $1
	AND entry_time >= $2`
	contactEvents, err := db.Conn.Query(query, userId, dt)
	if err != nil {
		return contactUsers, err
	}

	for contactEvents.Next() {
		var locationId int
		var entryTime string        // will never be null
		var exitTime sql.NullString // sometimes null
		err := contactEvents.Scan(
			&locationId,
			&entryTime,
			&exitTime,
		)
		if err != nil {
			return contactUsers, err
		}

		query = `SELECT user_id FROM entry_log
		WHERE user_id != $1
		AND location_id = $2
		AND (
			(entry_time <= $3 AND exit_time <= $4 AND exit_time >= $3)
			OR
			(entry_time <= $3 AND exit_time >= $4)
			OR
			(entry_time >= $3 AND exit_time <= $4)
			OR
			(entry_time >= $3 AND exit_time >= $4 AND entry_time <= $4)
		)`
		contactUserLogs, err := db.Conn.Query(
			query,
			userId,
			locationId,
			entryTime,
			exitTime,
		)
		if err != nil {
			return contactUsers, err
		}
		for contactUserLogs.Next() {
			var contactUserId string
			err := contactUserLogs.Scan(
				&contactUserId,
			)
			if err != nil {
				return contactUsers, err
			}
			contactUserIds[contactUserId] = true
		}
	}
	for id := range contactUserIds {
		user, err := db.GetUser(id)
		if err != nil {
			return contactUsers, err
		}
		contactUsers.Users = append(contactUsers.Users, *user)
	}
	return contactUsers, nil
}
