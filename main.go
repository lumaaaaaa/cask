package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	BaseURL      = "https://api.chaton.ai"
	Path         = "/chats/stream"
	Model        = "gpt-4o"
	MaxTokens    = 4096
	SystemPrompt = "You are ChatGPT, a large language model trained by OpenAI, based on the GPT-4o architecture. You are here to assist and provide information."
	Version      = "1.37.346"
)

var (
	client = &http.Client{}
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cask <prompt>")
		os.Exit(1)
	}

	args := strings.Join(os.Args[1:], " ")

	ask(args)
}
