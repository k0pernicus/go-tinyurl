package types

//Response structures and permits to encode the service response for any query
type Response struct {
	StatusCode int         `json:"status_code"`
	Response   interface{} `json:"response"`
}

// CreationRequest is a simple structure that contains all the informations
// to create a tiny URL
type CreationRequest struct {
	URL    string   `json:"url"`
	DeadIn Duration `json:"dead_in"`
}

// CreationResponse is the structure, returned by the service, when creating a tiny URL
type CreationResponse struct {
	ID      string  `json:"id,omitempty"`
	Message Message `json:"message"`
}

// ExistsResponse is a specific structure that handles the response for the exists handler
type ExistsResponse struct {
	Exists  bool    `json:"exists"`
	Message Message `json:"message,omitempty"`
}
