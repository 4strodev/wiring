package errors

import (
	"errors"
	"fmt"
)

type WiringError struct {
	error
}

func WrapError(err error) *WiringError {
	wiringError := new(WiringError)
	wiringError.error = err
	return wiringError
}

func NewError(message string) *WiringError {
	err := errors.New(message)
	return WrapError(err)
}

func Errorf(message string, arguments ...any) *WiringError {
	err := fmt.Errorf(message, arguments...)
	return WrapError(err)
}

func (err *WiringError) Error() string {
	return err.error.Error()
}
