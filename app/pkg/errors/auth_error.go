package errors

import "errors"

var (
	ErrEmailExist         = errors.New("email is already exist")
	ErrNameRequired       = errors.New("name is required")
	ErrEmailRequired      = errors.New("email is required")
	ErrPasswordRequired   = errors.New("password is required")
	ErrInvalidEmailFormat = errors.New("email format is invalid")
)
