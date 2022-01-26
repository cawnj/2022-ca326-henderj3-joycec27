package models

import (
	"fmt"
	"net/http"
)

type User struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	PhoneNumber   string `json:"phone_number"`
	CovidPositive bool   `json:"covid_positive"`
}

type UserList struct {
	Users []User `json:"users"`
}

func (u *User) Bind(r *http.Request) error {
	if u.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

func (*UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
