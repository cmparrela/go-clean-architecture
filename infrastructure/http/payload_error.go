package http

type PayloadError struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}
