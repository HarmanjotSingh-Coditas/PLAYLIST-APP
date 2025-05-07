package models

// Only used for Swagger generation
type ErrorMessage struct {
	Key          string `json:"key,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type ErrorAPIResponse struct {
	Message      []ErrorMessage `json:"errors,omitempty"`
	ErrorMessage string         `json:"error,omitempty"`
}
