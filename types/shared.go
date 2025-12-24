package types

// ErrorClientExceptionResponse represents a detailed error response from the API.
type ErrorClientExceptionResponse struct {
	StatusCode int    `json:"statusCode"`
	Timestamp  string `json:"timestamp"`
	Path       string `json:"path"`
	HashCode   string `json:"hashCode"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

// ErrorClientSimpleResponse represents a simple error response from the API.
type ErrorClientSimpleResponse struct {
	Message string `json:"message"`
}
