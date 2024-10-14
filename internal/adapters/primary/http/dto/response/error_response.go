package response

// ErrorResponse represents an error response
// @Description Error response structure
type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Type       string `json:"type"`
	RequestID  string `json:"request_id"`
	Message    string `json:"message"`
	Details    any    `json:"details,omitempty"`
}
