package internal

func (m *Error) Error() string {
	return m.Cause
}

// ParseError to parse proto error
func ParseError(err error) *Error {
	if o, ok := err.(*Error); ok {
		return o
	}
	return &Error{Cause: err.Error()}
}
