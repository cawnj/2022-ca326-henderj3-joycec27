package db

import (
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
	var entryTime sql.NullString
	query := `UPDATE entry_log
	SET exit_time = $2::timestamp
	WHERE entry_id = (
		SELECT max(entry_id)
		FROM entry_log
		WHERE user_id = $1
	)
	RETURNING entry_id, entry_time`
	err := db.Conn.QueryRow(
		query,
		entryLog.UserID,
		entryLog.ExitTime,
	).Scan(&id, &entryTime)
	if err != nil {
		return err
	}
	entryLog.EntryID = id
	entryLog.EntryTime = entryTime
	return nil
}

func (db Database) GetLatestEntryLog(userId int) (*models.EntryLog, error) {
	entryLog := &models.EntryLog{}
	query := `SELECT * FROM entry_log
	WHERE entry_id = (
		SELECT max(entry_id)
		FROM entry_log
		WHERE user_id = $1
	)`
	err := db.Conn.QueryRow(query, userId).Scan(
		&entryLog.EntryID,
		&entryLog.UserID,
		&entryLog.LocationID,
		&entryLog.EntryTime,
		&entryLog.ExitTime,
	)
	if err != nil {
		return nil, err
	}
	return entryLog, nil
}
