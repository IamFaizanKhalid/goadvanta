package goadvanta

import (
	"net/http"
	"time"
)

type sendSmsRequest struct {
	Apikey     string `json:"apikey"`
	PartnerID  string `json:"partnerID"`
	ShortCode  string `json:"shortcode"`
	Mobile     string `json:"mobile"`
	Message    string `json:"message"`
	TimeToSend string `json:"timeToSend;omitempty"`
}

type sendSmsResponses struct {
	Responses []sendSmsResponse `json:"responses"`
}

type sendSmsResponse struct {
	ResponseCode    int    `json:"response-code"`
	ResponseMessage Error  `json:"response-description"`
	Mobile          int    `json:"mobile"`
	MessageID       string `json:"messageid"`
	NetworkID       int    `json:"networkid"`
}

func (c client) SendSMS(phone, message string) (*MessageResponse, error) {
	return c.sendSMS(phone, message, nil)
}

func (c client) SendSMSLater(phone, message string, at time.Time) (*MessageResponse, error) {
	return c.sendSMS(phone, message, &at)
}

func (c client) sendSMS(phone, message string, schedule *time.Time) (*MessageResponse, error) {
	req := sendSmsRequest{
		Apikey:    c.apiKey,
		PartnerID: c.partnerID,
		ShortCode: c.shortCode,
		Mobile:    phone,
		Message:   message,
	}

	if schedule != nil {
		req.TimeToSend = schedule.Format("2006-01-02 15:04:05")
	}

	var resp sendSmsResponses
	err := c.request(http.MethodPost, "/sendsms/", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Responses[0].ResponseCode != http.StatusOK {
		return &MessageResponse{
			Error: resp.Responses[0].ResponseMessage,
		}, nil
	}

	return &MessageResponse{
		Success:   true,
		MessageID: resp.Responses[0].MessageID,
		NetworkID: resp.Responses[0].NetworkID,
	}, nil
}
