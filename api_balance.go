package goadvanta

import "net/http"

type getBalanceRequest struct {
	Apikey    string `json:"apikey"`
	PartnerID string `json:"partnerID"`
}

type getBalanceResponse struct {
	ResponseCode int    `json:"response-code"`
	Credit       string `json:"credit"`
	PartnerID    string `json:"partner-id"`
}

func (c client) GetBalance() (string, error) {
	req := getBalanceRequest{
		Apikey:    c.apiKey,
		PartnerID: c.partnerID,
	}

	var resp getBalanceResponse
	err := c.request(http.MethodPost, "/getbalance/", req, &resp)
	if err != nil {
		return "", err
	}

	return resp.Credit, nil
}
