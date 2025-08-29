package apperror

type errorText string

// Common errors.
const (
	InternalErrorText   errorText = "Internal error."
	BadRequestErrorText errorText = "Bad request."
)

// User errors.
const (
	LoginTakenErrorText errorText = "Login is already taken."
)
