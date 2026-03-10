package config

import "errors"

var (
	ErrKeyNotFound        = errors.New("key not found")
	ErrTypeMismatch       = errors.New("type mismatch")
	ErrInvalidEntry       = errors.New("invalid entry")
	ErrDotenvFileNotFound = errors.New("dotenv file not found")
)
