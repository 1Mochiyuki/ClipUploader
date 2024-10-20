package errs

import "errors"

type EmptyStringError struct {
	error
}

func NewEmptyStringError() EmptyStringError {
	return EmptyStringError{
		errors.New("string should not be empty"),
	}
}
