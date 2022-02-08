package db

import (
	"database/sql"
	"log"
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
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return entryLog, nil
	}
}

func (db Database) GetContactUsers(userId int) (*models.EntryLogList, error) {
	entryLogs := &models.EntryLogList{}
	query := `SELECT * FROM entry_log
	WHERE user_id = $1`
	contactEvents, err := db.Conn.Query(query, userId)
	if err != nil {
		return entryLogs, err
	}

	for contactEvents.Next() {
		var contactEvent models.EntryLog
		err := contactEvents.Scan(
			&contactEvent.EntryID,
			&contactEvent.UserID,
			&contactEvent.LocationID,
			&contactEvent.EntryTime,
			&contactEvent.ExitTime,
		)
		if err != nil {
			return entryLogs, err
		}

		log.Printf("%+v", contactEvent)
		query = `SELECT * FROM entry_log
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
			contactEvent.LocationID,
			contactEvent.EntryTime.String,
			contactEvent.ExitTime.String,
		)
		if err != nil {
			return entryLogs, err
		}
		for contactUserLogs.Next() {
			var contactUserLog models.EntryLog
			err := contactUserLogs.Scan(
				&contactUserLog.EntryID,
				&contactUserLog.UserID,
				&contactUserLog.LocationID,
				&contactUserLog.EntryTime,
				&contactUserLog.ExitTime,
			)
			if err != nil {
				return entryLogs, err
			}
			entryLogs.EntryLogs = append(entryLogs.EntryLogs, contactUserLog)
		}
	}
	return entryLogs, nil
}
