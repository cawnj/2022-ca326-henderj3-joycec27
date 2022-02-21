package models

import (
	"fmt"
	"net/http"
)

type UserRequest struct {
	UserID string `json:"user_id"`
}

func (u *UserRequest) Bind(r *http.Request) error {
	if u.UserID == "" {
		return fmt.Errorf("user_id is a required field")
	}
	return nil
}

func (*UserRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
