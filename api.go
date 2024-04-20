package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func ask(prompt string) {
	// init signature
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	prefix := []byte("POST" + ":" + Path + ":" + timestamp + "\n")

	// create request body
	var request RequestBody
	request.FunctionImageGen = true
	request.FunctionWebSearch = true
	request.MaxTokens = MaxTokens
	request.Messages = []struct {
		Content string `json:"content"`
		Role    string `json:"role"`
	}{
		{
			Content: SystemPrompt,
			Role:    "system",
		},
		{
			Content: prompt,
			Role:    "user",
		},
	}
	request.Model = Model
	body, err := json.Marshal(request)
	if err != nil {
		log.Fatal("Error marshalling request body: ", err)
	}

	// create signature
	signature := generateSignature(append(prefix, body...))

	// create authorization header
	secretAuthPrefix := []byte{252, 137, 185, 155, 127, 94, 106, 81, 69, 242, 189, 184, 26, 228, 174, 239}
	authorization := fmt.Sprintf("Bearer %s.%s", b64.StdEncoding.EncodeToString(secretAuthPrefix), signature)

	// create request
	req, _ := http.NewRequest("POST", BaseURL+Path, bytes.NewReader(body))
	req.Header.Add("Date", timestamp)
	req.Header.Add("Client-time-zone", "-04:00") // TODO: dynamic timezone
	req.Header.Add("Authorization", authorization)
	req.Header.Add("User-Agent", "ChatOn_Android/"+Version)
	req.Header.Add("Accept-Language", "en-US")
	req.Header.Add("X-Cl-Options", "hb")
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	// send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}
	defer resp.Body.Close()

	// read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	fmt.Println(string(respBody))

	// handle body
	content := ""
	split := strings.Split(string(respBody), "data: ")
	for _, s := range split {
		// ignore empty strings and end of transmission
		if len(s) != 0 && s != "[DONE]\x0a\x0a" {
			var chunk ResponseChunk
			err := json.Unmarshal([]byte(s), &chunk)
			if err != nil {
				log.Fatal("Error unmarshalling chunk: ", err)
			}
			if len(chunk.Choices) != 0 {
				content += chunk.Choices[0].Delta.Content
			}
		}
	}

	// print the trimmed chat message
	fmt.Println(strings.TrimSpace(content))
}

func generateSignature(toSign []byte) string {
	secretKey := []byte{14, 94, 79, 102, 38, 245, 11, 65, 100, 43, 115, 94, 15, 241, 14, 16, 66, 129, 248, 226, 98, 109, 235, 60, 62, 41, 78, 29, 72, 181, 47, 8}
	h := hmac.New(sha256.New, secretKey)
	h.Write(toSign)
	return b64.StdEncoding.EncodeToString(h.Sum(nil))
}
