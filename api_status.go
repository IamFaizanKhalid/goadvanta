package goadvanta

import (
	"errors"
	"net/http"
	"time"
)

type deliveryStatusRequest struct {
	Apikey    string `json:"apikey"`
	PartnerID string `json:"partnerID"`
	MessageID string `json:"messageID"`
}

type deliveryStatusResponse struct {
	ResponseCode        int    `json:"response-code"`
	MessageID           string `json:"message-id"`
	ResponseDescription string `json:"response-description"`
	DeliveryStatus      int    `json:"delivery-status"`
	DeliveryDescription string `json:"delivery-description"`
	DeliveryTat         string `json:"delivery-tat"`
	DeliveryNetworkId   int    `json:"delivery-networkid"`
	DeliveryTime        string `json:"delivery-time"`
}

func (c Client) GetDeliveryReport(messageID string) (*DeliveryStatusResponse, error) {
	req := deliveryStatusRequest{
		Apikey:    c.apiKey,
		PartnerID: c.partnerID,
		MessageID: messageID,
	}

	var resp deliveryStatusResponse
	err := c.request(http.MethodPost, "/getdlr/", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.ResponseCode != http.StatusOK {
		return nil, errors.New(resp.ResponseDescription)
	}

	deliveryTime, _ := time.Parse("2006-01-02 15:04:05", resp.DeliveryTime)
	tat, _ := time.Parse("15:04:05", resp.DeliveryTat)
	h, m, s := tat.Clock()
	tatDuration := h*3600 + m*60 + s

	return &DeliveryStatusResponse{
		MessageID:         resp.MessageID,
		NetworkId:         resp.DeliveryNetworkId,
		Status:            resp.DeliveryStatus,
		Description:       DeliveryDescription(resp.DeliveryDescription),
		TurnAroundSeconds: tatDuration,
		DeliveryTime:      deliveryTime,
	}, nil
}
