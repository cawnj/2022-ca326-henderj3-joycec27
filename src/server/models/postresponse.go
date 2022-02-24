package models

import (
	"net/http"
)

type PostResponse struct {
	StatusCode int    `json:"status_code"`
	StatusText string `json:"status_text"`
}

func (*PostResponse) Bind(r *http.Request) error {
	return nil
}

func (*PostResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
