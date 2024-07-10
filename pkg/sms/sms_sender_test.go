package sms

import (
	"fmt"
	"testing"
)

func TestSendSMS(t *testing.T) {
	sender := SmsSender{
		ApiKey:     "",
		ApiUser:    "",
		Recipients: []string{"+254745617596"},
		Message:    "",
		Sender:     "",
	}

	response, err := sender.SendSMS()
	if err != nil {
		t.Errorf("error occurred: %v", err)
	}

	if response.ErrorResponse.HasError {
		t.Errorf("error response received: %s", response.ErrorResponse.Message)
	}

	if len(response.SmsMessageData.Recipients) == 0 {
		t.Errorf("no recipients received in response")
	}

	if response.SmsMessageData.Message == "" {
		t.Errorf("empty message received in response")
	}
}

func (s *SmsSender) MockSendSMS() (SmsSenderResponse, error) {
	// Todo: Implement mock behavior
	return SmsSenderResponse{
		SmsMessageData: SmsMessageData{
			Message: "Mocked success message",
			Recipients: []Recipient{
				{
					Key:         "mock-recipient-key",
					Cost:        "0.05",
					SmsKey:      "mock-sms-key",
					MessageId:   "mock-message-id",
					MessagePart: 1,
					Number:      "+254745617596",
					Status:      "Success",
					StatusCode:  "200",
				},
			},
		},
	}, nil
}

func TestRetrySendSMS(t *testing.T) {
	sender := SmsSender{
		ApiKey:     "your-api-key",
		ApiUser:    "your-api-user",
		Recipients: []string{"+254745617596"},
		Message:    "Hello, this is a test message.",
		Sender:     "YourSenderID",
	}

	maxRetries := 5

	sender.SendSMS = sender.MockSendSMS

	response, err := sender.RetrySendSMS(maxRetries)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if response.SmsMessageData.Message != "Mocked success message" {
		t.Errorf("unexpected message received in response")
	}
}

func TestRetrySendSMS_Success(t *testing.T) {
	sender := SmsSender{
		ApiKey:     "your-api-key",
		ApiUser:    "your-api-user",
		Recipients: []string{"+254745617596"},
		Message:    "Hello, this is a test message.",
		Sender:     "YourSenderID",
	}

	maxRetries := 5

	retryCount := 0
	sender.SendSMS = func() (SmsSenderResponse, error) {
		if retryCount < 3 {
			retryCount++
			return SmsSenderResponse{}, fmt.Errorf("mocked error")
		}
		return SmsSenderResponse{
			SmsMessageData: SmsMessageData{
				Message: "Success!",
			},
		}, nil
	}

	response, err := sender.RetrySendSMS(maxRetries)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if response.SmsMessageData.Message != "Success!" {
		t.Errorf("unexpected message received in response")
	}
}

func TestRetrySendSMS_Failure(t *testing.T) {
	sender := SmsSender{
		ApiKey:     "your-api-key",
		ApiUser:    "your-api-user",
		Recipients: []string{"+254745617596"},
		Message:    "Hello, this is a test message.",
		Sender:     "YourSenderID",
	}

	maxRetries := 3

	sender.SendSMS = func() (SmsSenderResponse, error) {
		return SmsSenderResponse{}, fmt.Errorf("mocked error")
	}

	_, err := sender.RetrySendSMS(maxRetries)

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
