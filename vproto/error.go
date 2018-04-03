package vproto

import "fmt"

func (m *Error) Error() string {
	return m.Cause
}

func NewError(format string, a ...interface{}) *Error {
	return &Error{Cause: fmt.Sprintf(format, a...)}
}

// ParseError to parse proto error
func ParseError(err error) *Error {
	if o, ok := err.(*Error); ok {
		return o
	}
	return &Error{Cause: err.Error()}
}
