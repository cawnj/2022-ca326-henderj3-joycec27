package db

import (
	"database/sql"

	"sonic-server/models"
)

func (db Database) GetUser(userId int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT * FROM users
	WHERE user_id = $1`
	row := db.Conn.QueryRow(query, userId)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.CovidPositive)
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
		err := rows.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.CovidPositive)
		if err != nil {
			return users, err
		}
		users.Users = append(users.Users, user)
	}
	return users, nil
}

func (db Database) RegisterUser(registerReq *models.RegisterRequest) error {
	return nil
}
