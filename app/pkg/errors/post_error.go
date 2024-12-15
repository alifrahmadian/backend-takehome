package errors

import "errors"

var (
	ErrPostTitleRequired   = errors.New("title is required")
	ErrPostContentRequired = errors.New("content is required")
	ErrPostNotFound        = errors.New("post not found")
	ErrInvalidPostID       = errors.New("invalid post id")
)