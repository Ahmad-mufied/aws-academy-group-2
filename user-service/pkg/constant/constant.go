package constant

import "errors"

var (
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequest     = errors.New("bad request")
	ErrUnauthorized   = errors.New("unauthorized access")
	ErrForbidden      = errors.New("forbidden")
	ErrNotFound       = errors.New("resource not found")
	ErrConflict       = errors.New("conflict detected")

	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
	ErrInvalidUser  = errors.New("invalid user credentials")
)
