package sms_sender

import (
	"context"
	"github.com/MikeMwita/africastalking-go/pkg/sms"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	apiUser := os.Getenv("API_USER")

	client := sms.NewClient(&http.Client{}, apiKey, apiUser)
	sender := &sms.SmsSender{
		Client:     client,
		Recipients: []string{"+1234567890"},
		Message:    "Test message",
		Sender:     "YourSenderID",
		SmsKey:     "unique_sms_key",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := sender.RetrySendSMS(ctx, 3)
	if err != nil {
		log.Fatalf("Failed to send SMS: %v", err)
	}

	log.Printf("SMS Response: %+v", resp)
}
