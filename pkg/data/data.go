package data

import (
	"bytes"
	"encoding/json"
	"net/http"
)

/*
	sending mobile data package
	https://developers.africastalking.com/docs/data/overview
*/


type Recipient struct {
	PhoneNumber string `json:"phoneNumber"`
	Quantity    int    `json:"quantity"`
	Unit 	  string `json:"unit"`
	Valiidity   string `json:"validity"`
	IsPromo	 bool   `json:"isPromo"`
	Metadata   map[string]string `json:"metadata"`
}


type Request struct {
	UserName string `json:"username"`
	ProductName string `json:"productName"`
	Recipients []Recipient `json:"recipients"`
}

type Entry struct {
	PhoneNumber string `json:"phoneNumber"`
	Status string `json:"status"`
	TransactionId string `json:"transactionId"`
	Value string `json:"value"`
	Provider string `json:"provider"`
}


type Response struct {
	Entries []Entry `json:"entries"`
}


type Client struct {
	Username string
	ApiKey string
	Sandbox bool
	Client *http.Client	
}


func applyHeaders(req *http.Request, apiKey string){
	req.Header.Set("apiKey", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
}


func (c ( Client)) SendData(req Request) (*Response, error){
	url := "https://payments.africastalking.com/mobile/data/request"

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpClient := c.Client
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	applyHeaders(httpReq, c.ApiKey)

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()

	var resp Response
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}