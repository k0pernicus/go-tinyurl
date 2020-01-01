package types

import "time"

type Response struct {
	StatusCode int         `json:"status_code"`
	Response   interface{} `json:"response"`
}

// CreationRequest is a simple structure that contains all the informations
// to create a tiny URL
type CreationRequest struct {
	URL  string        `json:"url"`
	Time time.Duration `json:"duration"`
}

// CreationResponse is the structure, returned by the service, when creating a tiny URL
type CreationResponse struct {
	Message Message `json:"message"`
	ID      string  `json:"id,omitempty"`
}

type ExistsResponse struct {
	Exists  bool    `json:"exists"`
	Message Message `json:"message,omitempty"`
}
