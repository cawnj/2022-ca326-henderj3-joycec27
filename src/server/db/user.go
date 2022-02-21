package db

import (
	"database/sql"

	"sonic-server/models"
)

func (db Database) GetUser(userId string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT * FROM users
	WHERE user_id = $1`
	row := db.Conn.QueryRow(query, userId)
	err := row.Scan(&user.UserID, &user.ExpoToken)
	switch err {
	case sql.ErrNoRows:
		return nil, ErrNoMatch
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

func (db Database) GetAllUsers() (*models.UserList, error) {
	users := &models.UserList{}
	rows, err := db.Conn.Query("SELECT * FROM users ORDER BY user_id DESC")
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UserID, &user.ExpoToken)
		if err != nil {
			return users, err
		}
		users.Users = append(users.Users, user)
	}
	return users, nil
}

func (db Database) RegisterUser(userReq *models.User) error {
	user, err := db.GetUser(userReq.UserID)
	if err != nil && err != ErrNoMatch {
		return err
	}

	var query string
	switch {
	// user exists, update token
	case user != nil && userReq.ExpoToken != user.ExpoToken:
		query = `UPDATE users
		SET expo_token = $2
		WHERE user_id = $1`
	// user does not exist, insert new row
	case user == nil:
		query = `INSERT INTO users (user_id, expo_token)
		VALUES ($1, $2)`
	// user exists but has not changed, return
	default:
		return nil
	}

	_, err = db.Conn.Exec(query, userReq.UserID, userReq.ExpoToken)
	return err
}
