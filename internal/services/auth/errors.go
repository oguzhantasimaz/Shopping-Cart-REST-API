package auth_service

import "errors"

var (
	ErrUsernameRequired = errors.New("username is required")
	ErrPasswordRequired = errors.New("password is required")
	ErrTokenRequired    = errors.New("token is required")
)
