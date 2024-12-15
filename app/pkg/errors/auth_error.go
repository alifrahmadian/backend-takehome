package errors

import "errors"

var (
	ErrEmailExist            = errors.New("email is already exist")
	ErrNameRequired          = errors.New("name is required")
	ErrEmailRequired         = errors.New("email is required")
	ErrPasswordRequired      = errors.New("password is required")
	ErrInvalidEmailFormat    = errors.New("email format is invalid")
	ErrUserNotExist          = errors.New("this user is not exist")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrNoAuthorizationHeader = errors.New("authorization header is required")
	ErrInvalidTokenFormat    = errors.New("invalid token format")
	ErrTokenInvalid          = errors.New("invalid token")
	ErrTokenExpired          = errors.New("token expired")
)
