package types

import "fmt"

// ErrorCode represents error classification
type ErrorCode string

const (
	ErrorCodeNotFound       ErrorCode = "not_found"
	ErrorCodeUnauthorized   ErrorCode = "unauthorized"
	ErrorCodeForbidden      ErrorCode = "forbidden"
	ErrorCodeBadRequest     ErrorCode = "bad_request"
	ErrorCodeInternal       ErrorCode = "internal_error"
	ErrorCodeTimeout        ErrorCode = "timeout"
	ErrorCodeUnavailable    ErrorCode = "service_unavailable"
	ErrorCodeConflict       ErrorCode = "conflict"
)

// Error represents a structured error with code and context
type Error struct {
	Code    ErrorCode
	Message string
	Details map[string]interface{}
	Err     error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *Error) Unwrap() error {
	return e.Err
}

// NewError creates a new Error
func NewError(code ErrorCode, message string, err error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
		Details: make(map[string]interface{}),
	}
}

// WithDetails adds details to the error
func (e *Error) WithDetails(key string, value interface{}) *Error {
	e.Details[key] = value
	return e
}
