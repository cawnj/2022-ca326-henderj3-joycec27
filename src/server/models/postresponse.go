package models

import (
	"net/http"

	"github.com/go-chi/render"
)

type PostResponse struct {
	StatusCode int    `json:"status_code"`
	StatusText string `json:"status_text"`
}

func (*PostResponse) Bind(r *http.Request) error {
	return nil
}

func (p *PostResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, p.StatusCode)
	return nil
}
