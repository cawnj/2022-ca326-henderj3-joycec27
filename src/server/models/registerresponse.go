package models

import (
	"net/http"
)

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
