package sms

type SmsSender struct {
	Client     *Client
	Recipients []string `json:"recipients"`
	Message    string   `json:"message"`
	Sender     string   `json:"sender"`
	SmsKey     string   `json:"sms_key"`
}

type Recipient struct {
	Key         string `json:"key"`
	Cost        string `json:"cost"`
	SmsKey      string `json:"sms_key"`
	MessageId   string `json:"message_id"`
	MessagePart int    `json:"message_part"`
	Number      string `json:"number"`
	Status      string `json:"status"`
	StatusCode  string `json:"status_code"`
}

type SmsMessageData struct {
	Message    string      `json:"message"`
	Cost       string      `json:"cost"`
	Recipients []Recipient `json:"recipients"`
}
type ErrorResponse struct {
	HasError bool   `json:"has_error"`
	Message  string `json:"message"`
}

type SmsSenderResponse struct {
	ErrorResponse  ErrorResponse  `json:"error_response"`
	SmsMessageData SmsMessageData `json:"sms_message_data"`
}
