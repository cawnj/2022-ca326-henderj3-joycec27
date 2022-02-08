package models

import (
	"fmt"
	"net/http"
)

type TraceRequest struct {
	UserID int `json:"user_id"`
}

func (t *TraceRequest) Bind(r *http.Request) error {
	if t.UserID == 0 {
		return fmt.Errorf("user_id is a required field")
	}
	return nil
}

func (*TraceRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
