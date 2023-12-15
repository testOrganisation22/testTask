package error

import "errors"

var (
	ServerError      = errors.New("server error")
	IncorrectHash    = errors.New("incorrect hash")
	IncorrectAddress = errors.New("incorrect address")
)
