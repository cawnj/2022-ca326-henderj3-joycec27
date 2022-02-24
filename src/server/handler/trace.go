package handler

import (
	"fmt"
	"net/http"

	"sonic-server/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

func trace(router chi.Router) {
	router.Post("/", contactTrace)
}

func contactTrace(w http.ResponseWriter, r *http.Request) {
	userReq := &models.UserRequest{}
	if err := render.Bind(r, userReq); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	users, err := dbInstance.GetContactUsers(userReq.UserID)
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := notifyCloseContacts(users); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
	response := &models.PostResponse{
		StatusCode: 201,
		StatusText: fmt.Sprintf("notified %d close contact(s)", len(users.Users)),
	}
	if err := render.Render(w, r, response); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func notifyCloseContacts(users *models.UserList) error {
	tokens := []expo.ExponentPushToken{}
	for _, user := range users.Users {
		pushToken, err := expo.NewExponentPushToken(user.ExpoToken)
		if err != nil {
			return err
		}
		tokens = append(tokens, pushToken)
	}

	client := expo.NewPushClient(nil)
	response, err := client.Publish(
		&expo.PushMessage{
			To:       tokens,
			Body:     "You are a close contact!",
			Sound:    "default",
			Title:    "Sonic",
			Priority: expo.DefaultPriority,
		},
	)
	if err != nil {
		return err
	}
	if response.ValidateResponse() != nil {
		return fmt.Errorf("%s failed", response.PushMessage.To)
	}

	return err
}
