package store

import (
	"errors"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateAccount   = errors.New("account already exists")
)

const internalErrorMessage = "internal error"

type InternalError struct {
	Err error
}

func NewInternalErr(err error) error {
	return err
}

func (*InternalError) Error() string {
	return internalErrorMessage
}

func (self *InternalError) InternalError() string {
	return self.Err.Error()
}
