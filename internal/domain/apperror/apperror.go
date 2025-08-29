package apperror

type AppError struct {
	Code    errorCode `json:"code,omitempty"`
	Error   errorText `json:"error"`
	Message string    `json:"message,omitempty"`
}
