package types

// Status contains a specific string that specifies if the creation has failed or not
type Status string

const (
	// CreatedStatus is returned if the tiny URL has correctly been created
	CreatedStatus Status = "Created"
	// ErrorStatus is returned if the tiny URL has not been created
	ErrorStatus Status = "Error"
)
