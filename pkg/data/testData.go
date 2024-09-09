package data

import (
	"os"
	"testing"
)



func TestSendData(t *testing.T) {
	client := &Client{
		ApiKey: os.Getenv("API_KEY"),
		Username: os.Getenv("USERNAME"),
		Sandbox: true,
	}

	request := &Request{
		UserName: "sandbox",
		ProductName: "data",
		Recipients: []Recipient{
			{PhoneNumber: "+254757387606", Quantity: 1, Unit: "MB", Valiidity: "1", IsPromo: false, Metadata: map[string]string{"name": "John Doe"}},
		},
	}
	resp, err := client.SendData(*request)
	if err != nil {
		t.Errorf("Error sending data: %v", err)
	}
	status := resp.Entries[0].Status
	if status != "Success" {
		t.Errorf("Expected status to be Success, got %v", status)
	}
}