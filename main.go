package main

import (
	"net/http"
)

const (
	BaseURL      = "https://api.chaton.ai"
	Path         = "/chats/stream"
	Model        = "gpt-4-0125-preview"
	MaxTokens    = 4096
	SystemPrompt = "You are ChatGPT, a large language model trained by OpenAI, based on the GPT-4 architecture. You are here to assist and provide information."
	Version      = "1.37.346"
)

var (
	client = &http.Client{}
)

func main() {
	ask("Is this thing on? Who are you?")
}
