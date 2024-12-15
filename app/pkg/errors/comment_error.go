package errors

import "errors"

var (
	ErrCommentContentRequired = errors.New("comment content is required")
)
