package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)


type Messenger interface {
	SendMessage(webhookUrl string, msg string) error
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Data format for Teams message card
// The field tag results in the field being marshalled as "text" instead of "Text"
type TeamsMessageCard struct {
	Text string `json:"text"`
}

type TeamsMessenger struct {
	Client HTTPClient
}

// (t TeamsMessenger) is a receiver of the SendMessage method
func (t TeamsMessenger) SendMessage(webhookUrl string, msg TeamsMessageCard) error {
    jsonMsg, err := json.Marshal(msg)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(jsonMsg))
    if err != nil {
        return err
    }
    req.Header.Set("Content-Type", "application/json")

    _, err = t.Client.Do(req)
    return err
}
