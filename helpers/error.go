package helpers

import "errors"

const (
	StrUnprocessableEntity = "Unprocessable Entity"
	StrInternalServerError = "Internal Server Error"
)

var (
	ErrDuplicateEmail = errors.New("email already registered")
	ErrEmailNotFound  = errors.New("email not found")
	ErrWrongPassword  = errors.New("wrong password")
	ErrInvalidToken   = errors.New("invalid token")
	ErrUnauthorized   = errors.New("unauthorized")
)