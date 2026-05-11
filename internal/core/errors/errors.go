package core_errors

import "errors"

var (
	ErrNotFound        = errors.New("Not found")
	ErrInvalidArgument = errors.New("Invalid argument")
)
