package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	webhookUrl := os.Getenv("WEBHOOK_URL")
	if webhookUrl == "" {
		log.Fatal("WEBHOOK_URL not set in environment")
	}

	msg := TeamsMessageCard{
		Text: "Hello from Go app!",
	}

	messenger := TeamsMessenger{
		Client: http.DefaultClient,
	}
	err := messenger.SendMessage(webhookUrl, msg)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
}
