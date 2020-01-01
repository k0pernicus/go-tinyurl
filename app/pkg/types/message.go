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
	// URLDoesNotExists is sent back if the requested tiny URL id does not exists
	URLDoesNotExists Message = "URL does not exists"
)
