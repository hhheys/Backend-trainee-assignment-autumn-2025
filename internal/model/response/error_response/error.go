// Package response provides types for structured API responses,
// including error handling structures used to standardize error messages
// returned by the server.
package response

// ErrorResponse represents the response structure for error responses.
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail represents the details of an error.
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse creates a new ErrorResponse instance.
func NewErrorResponse(code, message string) ErrorResponse {
	return ErrorResponse{
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	}
}
