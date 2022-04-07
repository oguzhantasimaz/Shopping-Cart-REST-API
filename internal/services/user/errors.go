package user_service

import "errors"

var (
	ErrUserNameRequired = errors.New("username is required")
	ErrPasswordRequired = errors.New("password is required")
	ErrUserIDRequired   = errors.New("user id is required")
)
