package sms

import "net/http"

const (
	DefaultAPIURL = "https://api.africastalking.com/version1/messaging"
	sandboxAPIURL = "https://api.sandbox.africastalking.com/version1/messaging"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	apiURL  string
	apiKey  string
	apiUser string
	client  Doer
}

func NewClient(client Doer, apiKey, apiUser string) *Client {
	return &Client{
		apiURL:  DefaultAPIURL,
		apiKey:  apiKey,
		apiUser: apiUser,
		client:  client,
	}
}
