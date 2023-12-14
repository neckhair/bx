package bexio

import "errors"

var (
	UnauthorizedError = errors.New("Unauthorized")
	NotFoundError     = errors.New("Not found")
)
