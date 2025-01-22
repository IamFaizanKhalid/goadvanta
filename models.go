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

type DeliveryDescription string

const (
	StatusDeliveredToTerminal   DeliveryDescription = "DeliveredToTerminal"
	StatusAbsentSubscriber      DeliveryDescription = "AbsentSubscriber"
	StatusDeliveryImpossible    DeliveryDescription = "Delivery Impossible"
	StatusSenderNameBlacklisted DeliveryDescription = "SenderName Blacklisted"
	StatusInvalidSourceAddress  DeliveryDescription = "Invalid Source Address"
)

func (d DeliveryDescription) IsSuccess() bool {
	return d == StatusDeliveredToTerminal
}

func (d DeliveryDescription) IsFailure() bool {
	return d != StatusDeliveredToTerminal
}

func (d DeliveryDescription) String() string {
	return string(d)
}

// Description based on https://sms.api.swifttdial.com/docs/SMS/DLR%20Status/
func (d DeliveryDescription) Description() string {
	switch d {
	case StatusDeliveredToTerminal:
		return "Delivered to device"
	case StatusAbsentSubscriber:
		return "Subscriber has not been able to connect to the network for the last 24 hrs. In some cases (airtel) customers have to lock their phone to 3G only to be able to receive messages. This is an issue specific to them"
	case StatusDeliveryImpossible:
		return "Number has been ported to another telco or has temporarily been deactivated by SP. Or the number has been inactive for a long period of time. In some cases the telco has blocked delivery due to violation of policy. ie. Sending marketing message beyond 6PM and earlier than 7am"
	case StatusSenderNameBlacklisted:
		return "The customer has explicitly blocked messages from the sender id. or opted out of receiving promotional messages"
	case StatusInvalidSourceAddress:
		return "The sender ID does not exist on the telco"
	}
	return ""
}

type DeliveryStatusResponse struct {
	MessageID         string
	NetworkId         int
	Status            int
	Description       DeliveryDescription
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
