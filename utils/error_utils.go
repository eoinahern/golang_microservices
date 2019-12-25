package utils

//ApplicationError : an error
type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
}
