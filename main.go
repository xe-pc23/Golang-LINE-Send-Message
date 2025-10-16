package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Request struct {
	To       string    `json:"to"`
	Messages []Message `json:"messages"`
}

func main() {
	godotenv.Load()

	token := os.Getenv("CHANNEL_ACCESS_TOKEN")
	userID := os.Getenv("USER_ID")

	req := Request{
		To:       userID,
		Messages: []Message{{Type: "text", Text: "テストメッセージです。"}},
	}

	jsonData, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", "https://api.line.me/v2/bot/message/push", bytes.NewBuffer(jsonData))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+token)

	http.DefaultClient.Do(httpReq)
}
