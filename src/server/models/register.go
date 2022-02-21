package models

import (
	"fmt"
	"net/http"
)

type RegisterRequest struct {
	UserID    string `json:"user_id"`
	ExpoToken string `json:"expo_token"`
}

func (req *RegisterRequest) Bind(r *http.Request) error {
	if req.UserID == "" {
		return fmt.Errorf("user_id is a required field")
	}
	if req.ExpoToken == "" {
		return fmt.Errorf("token is a required field")
	}
	return nil
}

func (*RegisterRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type RegisterResponse struct {
	StatusCode int    `json:"status_code"`
	StatusText string `json:"status_text"`
}

func (*RegisterResponse) Bind(r *http.Request) error {
	return nil
}

func (*RegisterResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
