package types

// Message is a type that contains a specific message when receiving a message from the service
type Message string

const (
	// OK seems... fine
	OK Message = "OK"
	// InternalError signifies that something gone wrong
	InternalError Message = "Internal Error"
	// CannotDecodeMessage if the body message is not correctly formated
	CannotDecodeMessage Message = "Cannot decode message"
	// CannotGenerateNewUUID if the UUID generation from google library returned an error
	CannotGenerateNewUUID Message = "Cannot generate new UUID"
	// CannotParseUUID if we can't parse UUID based on ID sent by user
	CannotParseUUID Message = "Cannot parse UUID"
	// URLDoesNotExists is sent back if the requested tiny URL id does not exists
	URLDoesNotExists Message = "URL does not exists"
)
