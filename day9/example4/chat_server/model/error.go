package model

import "errors"

var (
	ErrUserExist     = errors.New("user exist")
	ErrUserNotExist  = errors.New("user not exist")
	ErrInvalidPasswd = errors.New("passwd or username not right")
	ErrInvalidParams = errors.New("invalid params")
)
