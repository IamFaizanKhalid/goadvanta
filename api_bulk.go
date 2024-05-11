package goadvanta

import (
	"errors"
	"net/http"
)

type sendBulkSmsRequest struct {
	Count   int           `json:"count"`
	SmsList []smsListItem `json:"smslist"`
}

type smsListItem struct {
	Apikey      string `json:"apikey"`
	PartnerID   string `json:"partnerID"`
	ShortCode   string `json:"shortcode"`
	PassType    string `json:"pass_type"`
	ClientSmsId int    `json:"clientsmsid"`
	Mobile      string `json:"mobile"`
	Message     string `json:"message"`
}

type sendBulkSmsResponses struct {
	Responses []sendBulkSmsResponse `json:"responses"`
}

type sendBulkSmsResponse struct {
	ResponseCode    int         `json:"response-code"`
	ResponseMessage Error       `json:"response-description"`
	Mobile          interface{} `json:"mobile"`
	MessageID       string      `json:"messageid"`
	ClientSmsId     int         `json:"clientsmsid"`
	NetworkID       int         `json:"networkid"`
}

// SendBulkSMS can now send up to 20 bulk messages in one single call
func (c client) SendBulkSMS(messages []MessageRequest) ([]MessageResponse, error) {
	if len(messages) > 20 {
		return nil, errors.New("you can only send up to 20 messages at a time")
	}

	req := sendBulkSmsRequest{
		Count:   len(messages),
		SmsList: []smsListItem{},
	}

	for _, msg := range messages {
		req.SmsList = append(req.SmsList, smsListItem{
			PartnerID:   c.partnerID,
			Apikey:      c.apiKey,
			PassType:    "plain",
			ClientSmsId: msg.ClientID,
			Mobile:      msg.Phone,
			Message:     msg.Text,
			ShortCode:   c.shortCode,
		})
	}

	var resp sendBulkSmsResponses
	err := c.request(http.MethodPost, "/sendbulk/", req, &resp)
	if err != nil {
		return nil, err
	}

	var messagesResponse []MessageResponse
	for _, r := range resp.Responses {
		if r.ResponseCode != http.StatusOK {
			messagesResponse = append(messagesResponse, MessageResponse{
				Error: r.ResponseMessage,
			})
		} else {
			messagesResponse = append(messagesResponse, MessageResponse{
				Success:   true,
				ClientID:  r.ClientSmsId,
				MessageID: r.MessageID,
				NetworkID: r.NetworkID,
			})
		}
	}

	return messagesResponse, nil
}
