package types

import "errors"

func SizeOfStdError(d error) int {
	if d != nil {
		err := &Error{ErrorMessage: d.Error()}
		return err.Size()
	}
	return 0
}

func StdErrorMarshal(d error) ([]byte, error) {
	size := SizeOfStdError(d)
	buf := make([]byte, size)
	_, err := StdErrorMarshalTo(d, buf)
	return buf, err
}

func StdErrorMarshalTo(d error, data []byte) (int, error) {
	if d != nil {
		err := &Error{ErrorMessage: d.Error()}
		return err.MarshalTo(data)
	}
	return 0, nil
}

func StdErrorUnmarshal(d *error, data []byte) error {
	if len(data) > 0 {
		err := &Error{}
		if err := err.Unmarshal(data); err != nil {
			return err
		}
		if len(err.ErrorMessage) > 0 {
			*d = errors.New(err.ErrorMessage)
		}
	}
	return nil
}
