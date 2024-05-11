package goadvanta

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

const baseUrl = "https://quicksms.advantasms.com/api/services"

type client struct {
	apiKey    string
	partnerID string
	shortCode string
}

func NewClient(apiKey, partnerID, shortCode string) *client {
	return &client{
		apiKey:    apiKey,
		partnerID: partnerID,
		shortCode: shortCode,
	}
}

func (c client) request(method, path string, req interface{}, resp interface{}) error {
	var buf io.Reader
	if req != nil {
		reqBody, err := json.Marshal(req)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(reqBody)
	}

	httpReq, err := http.NewRequest(method, baseUrl+path, buf)
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	apiResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer apiResp.Body.Close()

	respBody, err := io.ReadAll(apiResp.Body)
	if err != nil {
		return err
	}

	log.Println(string(respBody))

	if strings.Contains(string(respBody), "errors") {
		var errResp errorResponse
		if err = json.Unmarshal(respBody, &errResp); err != nil {
			return err
		}

		return errResp
	}

	if resp == nil {
		return nil
	}

	return json.Unmarshal(respBody, resp)
}

type errorResponse struct {
	ResponseCode    int   `json:"response-code"`
	ResponseMessage Error `json:"response-description"`
	Errors          map[string]map[string]string
}

func (e errorResponse) Error() string {
	var errMsg string
	for _, v := range e.Errors {
		for _, vv := range v {
			errMsg += vv + ". "
		}
	}

	return errMsg
}
