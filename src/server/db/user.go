package db

import (
	"database/sql"

	"sonic-server/models"
)

func (db Database) GetUser(userId int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT * FROM users
	WHERE user_id = $1`
	err := db.Conn.QueryRow(
		query,
		userId,
	).Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.CovidPositive)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return user, nil
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

func (db Database) UpdateCovidPositive(userId int, value bool) (*models.User, error) {
	user := &models.User{}
	query := `UPDATE users
	SET covid_positive = $2
	WHERE user_id = $1`
	err := db.Conn.QueryRow(
		query,
		userId,
		value,
	).Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.CovidPositive)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return user, nil
	}
}
