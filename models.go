package goadvanta

import "time"

type MessageRequest struct {
	ClientID int
	Phone    string
	Text     string
}

type MessageResponse struct {
	ClientID  int
	Success   bool
	Error     Error
	MessageID string
	NetworkID int
}

type DeliveryStatusResponse struct {
	MessageID         string
	NetworkId         int
	Status            int
	Description       string
	TurnAroundSeconds int
	DeliveryTime      time.Time
}

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrInvalidSenderID        Error = "Invalid sender id"
	ErrNetworkNotAllowed      Error = "Network not allowed"
	ErrInvalidMobileNumber    Error = "Invalid mobile number"
	ErrLowBulkCredits         Error = "Low bulk credits"
	ErrSystemError            Error = "Failed. System error"
	ErrInvalidCredentials     Error = "Invalid credentials"
	ErrNoDeliveryReport       Error = "No Delivery Report"
	ErrUnsupportedDataType    Error = "unsupported data type"
	ErrUnsupportedRequestType Error = "unsupported request type"
	ErrInternalError          Error = "Internal Error. Try again after 5 minutes"
	ErrNoPartnerIDSet         Error = "No Partner ID is Set"
	ErrNoAPIKeyProvided       Error = "No API KEY Provided"
	ErrDetailsNotFound        Error = "Details Not Found"
	ErrNoDlr                  Error = "No dlr"
)
