package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (s *SmsSender) SendSMS(ctx context.Context) (SmsSenderResponse, error) {
	form := url.Values{}
	form.Add("username", s.Client.apiUser)
	form.Add("to", strings.Join(s.Recipients, ","))
	form.Add("message", s.Message)
	form.Add("from", s.Sender)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.Client.apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		return SmsSenderResponse{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apiKey", s.Client.apiKey)

	res, err := s.Client.client.Do(req)
	if err != nil {
		return SmsSenderResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return SmsSenderResponse{
			ErrorResponse: ErrorResponse{
				HasError: true,
				Message:  "Message not sent",
			},
		}, fmt.Errorf("status code: %d", res.StatusCode)
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return SmsSenderResponse{}, err
	}

	smsMessageData := data["SMSMessageData"].(map[string]interface{})
	message := smsMessageData["Message"].(string)
	cost := ""
	for _, word := range strings.Split(message, " ") {
		cost = word
	}

	recipientsData := smsMessageData["Recipients"].([]interface{})
	recipients := make([]Recipient, 0)

	for _, recipient := range recipientsData {
		recipientData := recipient.(map[string]interface{})

		rct := Recipient{
			Key:         uuid.New().String(),
			Cost:        recipientData["cost"].(string),
			SmsKey:      s.SmsKey,
			MessageId:   recipientData["messageId"].(string),
			MessagePart: int(recipientData["messageParts"].(float64)),
			Number:      recipientData["number"].(string),
			Status:      recipientData["status"].(string),
			StatusCode:  fmt.Sprintf("%v", recipientData["statusCode"]),
		}

		recipients = append(recipients, rct)
	}

	return SmsSenderResponse{
		ErrorResponse: ErrorResponse{
			HasError: false,
		},
		SmsMessageData: SmsMessageData{
			Message:    message,
			Cost:       cost,
			Recipients: recipients,
		},
	}, nil
}

func (s *SmsSender) RetrySendSMS(ctx context.Context, maxRetries int) (SmsSenderResponse, error) {
	for retry := 0; retry < maxRetries; retry++ {
		response, err := s.SendSMS(ctx)
		if err == nil {
			return response, nil
		}

		delay := time.Duration(1<<uint(retry)) * time.Second
		jitter := time.Duration(rand.Intn(int(delay))) * time.Millisecond
		waitTime := delay + jitter

		time.Sleep(waitTime)
	}
	return SmsSenderResponse{}, fmt.Errorf("max retries reached")
}
