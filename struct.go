package main

type RequestBody struct {
	FunctionImageGen  bool      `json:"function_image_gen"`
	FunctionWebSearch bool      `json:"function_web_search"`
	MaxTokens         int       `json:"max_tokens"`
	Messages          []Message `json:"messages"`
	Model             string    `json:"model"`
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type ResponseChunk struct {
	Id                string      `json:"id"`
	Object            string      `json:"object"`
	Created           int         `json:"created"`
	Model             string      `json:"model"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
	Choices           []struct {
		Index int `json:"index"`
		Delta struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"delta"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason interface{} `json:"finish_reason"`
	} `json:"choices"`
}

type CommandResponse struct {
	Commands []string `json:"commands"`
	Message  string   `json:"message"`
}
