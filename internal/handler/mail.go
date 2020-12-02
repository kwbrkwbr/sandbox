package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"sandbox/infra"
)

type MailRequest struct {
	From   string `json:"from,omitempty" validate:"required,email"`
	To     string `json:"to,omitempty" validate:"required,email"`
	Title  string `json:"title,omitempty" validate:"required"`
	Body   string `json:"body,omitempty" validate:"required"`
	Params string `json:"params,omitempty" validate:"required"`
}

func Mail(c echo.Context) error {
	r := new(MailRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	client := infra.NewPubsubClient("portfolio-297301")
	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	if err := client.SimplePublish("eventarc-us-central1-sendgrid-trigger-064", b); err != nil {
		return err
	}

	return c.String(http.StatusAccepted, "accepted")
}
