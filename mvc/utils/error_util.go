package utils

// ApplicationError contains the structure and details of our Errors
type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}
