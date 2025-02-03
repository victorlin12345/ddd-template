package presentation

import "fmt"

type ErrorCode int

const (
	// System Errors (1-999)
	ErrInternal     ErrorCode = 1
	ErrUnauthorized ErrorCode = 2
	ErrForbidden    ErrorCode = 3
	ErrNotFound     ErrorCode = 4
	ErrValidation   ErrorCode = 5

	// Domain Errors (1000-1999)
	ErrOrderNotFound     ErrorCode = 1000
	ErrOrderAlreadyExist ErrorCode = 1001
	ErrInvalidOrderState ErrorCode = 1002

	// Infrastructure Errors (2000-2999)
	ErrDatabase ErrorCode = 2000
	ErrRedis    ErrorCode = 2001
	ErrRPC      ErrorCode = 2002
)

type AppError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("error code: %d, message: %s, details: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("error code: %d, message: %s", e.Code, e.Message)
}

func New(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func Wrap(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func Is(err error, code ErrorCode) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == code
	}
	return false
}
