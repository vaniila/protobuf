package types

import "fmt"

// ProtoError equivalents to Error interface
type ProtoError interface {
	error

	// Returns the short phrase depicting the classification of the error.
	Code() string

	// Returns the error details message.
	Message() string

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErr() error
}

// Code returns error code
func (o Error) Code() string {
	return o.ErrorCode
}

// Message return error message
func (o Error) Message() string {
	return o.ErrorMessage
}

func (o Error) Error() string {
	return fmt.Sprintf("ErrorCode: %s, ErrorMessage: %s", o.ErrorCode, o.ErrorMessage)
}

// OrigErr returns original error
func (o Error) OrigErr() error {
	return fmt.Errorf(o.Error())
}

// NewError creates error
func NewError(errorCode string, errorMessage string) *Error {
	return &Error{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
}
