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
	// CannotCreateQRCode is sent back if the user requested a qr code but generation failed
	CannotCreateQRCode Message = "Cannot create QR code"
	// CannotInsertRecord is sent back if we can't insert a given record in SQLite DB
	CannotInsertRecord Message = "Cannot insert record"
	// CannotRetrieveRecord is sent back if we can't retrieve a given record from a SQLite DB
	CannotRetrieveRecord Message = "Cannot retrieve record"
)
