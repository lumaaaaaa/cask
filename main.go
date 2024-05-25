package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	BaseURL         = "https://api.chaton.ai"
	Path            = "/chats/stream"
	DefaultModel    = "gpt-4o"
	MaxTokens       = 4096
	SystemPrompt    = "You are ChatGPT, a large language model trained by OpenAI, based on the GPT-4o architecture. You are here to assist and provide information."
	CMDSystemPrompt = "You are ChatGPT, a large language model trained by OpenAI, based on the GPT-4o architecture. You are here to provide terminal commands to solve a problem based on the given query.\n\nYou will be provided in a system message containing the user's operating system and current working directory info.\n\nYou are to provide responses in JSON format, containing a array of terminal commands, with the following structure: {\"commands\":[\"<terminal command 1>\", ...]}\n\nPlease provide a response in the specified format, skip any additional formatting, pure JSON only. Avoid using the backtick character (`) in your responses.\n\n---"
	Version         = "1.37.346"
)

var (
	client = &http.Client{}
)

func printHelp() {
	fmt.Println("// cask - AI-powered chat interface -", Version)
	fmt.Println("usage: cask [args] <prompt> ")
	fmt.Println("arguments:")
	fmt.Println("\t-h, --help\t\tshow this help message and exit")
	fmt.Println("\t-v, --version\t\tshow version information and exit")
	fmt.Println("\t-c, --cmd\t\tenable the command mode")
	fmt.Println("\t-m, --model <model name>\t\tset the model to use")
}

func main() {
	mode := "default"
	model := DefaultModel
	promptStart := 0

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-v", "--version":
			fmt.Println(Version)
			os.Exit(0)

		case "-h", "--help":
			printHelp()
			os.Exit(0)

		case "-c", "--cmd":
			mode = "cmd"
			break

		case "-m", "--model":
			i++
			if model != DefaultModel {
				printHelp()
				os.Exit(1)
			}
			model = os.Args[i]
			break

		default:
			promptStart = i
			break
		}

		if promptStart != 0 {
			break
		}
	}

	if len(os.Args) < promptStart+1 {
		printHelp()
		os.Exit(1)
	}

	args := strings.Join(os.Args[promptStart:], " ")

	handleChat(args, model, mode)
}
