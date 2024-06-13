package ai_integration

import (
	"net/http"
)

type Requester interface {
	Send(client *http.Client) (string, error)
}

func NewRequester(userPrompt string) Requester {
	return &developmentRequest{
		Model: "mixtral-8x7b-32768",
		Messages: []requestMessage{
			{"user", userPrompt},
		},
	}
}
