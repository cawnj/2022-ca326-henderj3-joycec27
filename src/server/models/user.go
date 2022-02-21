package models

import (
	"fmt"
	"net/http"
)

type User struct {
	UserID    string `json:"user_id"`
	ExpoToken string `json:"expo_token"`
}

type UserList struct {
	Users []User `json:"users"`
}

func (u *User) Bind(r *http.Request) error {
	if u.UserID == "" {
		return fmt.Errorf("user_id is a required field")
	}
	if u.ExpoToken == "" {
		return fmt.Errorf("expo_token is a required field")
	}
	return nil
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
