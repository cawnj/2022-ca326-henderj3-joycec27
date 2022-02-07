package db

import (
	"sonic-server/models"
)

func (db Database) GetAllEntryLogs() (*models.EntryLogList, error) {
	entryLogs := &models.EntryLogList{}
	rows, err := db.Conn.Query("SELECT * FROM entry_log ORDER BY entry_id DESC")
	if err != nil {
		return entryLogs, err
	}
	for rows.Next() {
		var entryLog models.EntryLog
		err := rows.Scan(&entryLog.UserID, &entryLog.LocationID, &entryLog.EntryTime, &entryLog.ExitTime)
		if err != nil {
			return entryLogs, err
		}
		entryLogs.EntryLogs = append(entryLogs.EntryLogs, entryLog)
	}
	return entryLogs, nil
}

func (db Database) AddEntryLog(entryLog *models.EntryLog) error {
	var id int
	query := `INSERT INTO entry_log (user_id, location_id, entry_time, exit_time)
	VALUES ($1, $2, $3::timestamp, $4::timestamp)
	RETURNING entry_id`
	err := db.Conn.QueryRow(
		query,
		entryLog.UserID,
		entryLog.LocationID,
		entryLog.EntryTime,
		entryLog.ExitTime,
	).Scan(&id)
	if err != nil {
		return err
	}
	entryLog.EntryID = id
	return nil
}

func (db Database) UpdateEntryLog(entryLog *models.EntryLog) error {
	var id int
	query := `UPDATE entry_log
	SET exit_time = $2::timestamp
	WHERE user_id = $1
	ORDER BY entry_time DESC
	LIMIT 1
	RETURNING entry_id`
	err := db.Conn.QueryRow(
		query,
		entryLog.UserID,
		entryLog.ExitTime,
	).Scan(&id)
	if err != nil {
		return err
	}
	entryLog.EntryID = id
	return nil
}
