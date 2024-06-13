package ai_integration

import (
	_ "embed"
	"fmt"
	"net/http"
	"regexp"
)

// Generator used by the app with ChatGPT
var Generator Service = &generator{}

type Service interface {
	Generate(input string) (response string, err error)
}

var (
	//go:embed user_prompt.txt
	userPrompt string
)

type generator struct{}

func (g *generator) Generate(input string) (response string, err error) {
	response, err = NewRequester(fmt.Sprintf(userPrompt, input)).Send(http.DefaultClient)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`https?://[^\s]+`)
	return re.FindString(response), nil
}
