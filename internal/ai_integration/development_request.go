package ai_integration

import (
	"bytes"
	"encoding/json"
	"io"

	"net/http"

	"github.com/leapkit/core/envor"
)

type requestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type developmentRequest struct {
	Model    string           `json:"model"`
	Messages []requestMessage `json:"messages"`
}

func (req *developmentRequest) Send(client *http.Client) (string, error) {
	request := map[string]interface{}{
		"model":    req.Model,
		"messages": req.Messages,
	}

	bson, _ := json.Marshal(request)
	r, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(bson))
	if err != nil {
		return "", err
	}

	r.Header.Add("Authorization", "Bearer "+envor.Get("AI_KEY", ""))
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}

	jsonresp, _ := io.ReadAll(resp.Body)
	var res developmentResponse
	err = json.Unmarshal(jsonresp, &res)
	if err != nil {
		return "", err
	}

	if len(res.Choices) == 0 {
		return "", nil
	}

	return res.Choices[0].Message.Content, nil
}
