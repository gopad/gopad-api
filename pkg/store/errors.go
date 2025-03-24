package store

import (
	"errors"
)

var (
	// ErrWrongCredentials is returned when credentials are wrong.
	ErrWrongCredentials = errors.New("wrong credentials provided")

	// ErrAlreadyAssigned defines the error if relation is already assigned.
	ErrAlreadyAssigned = errors.New("user pack already exists")

	// ErrNotAssigned defines the error if relation is not assigned.
	ErrNotAssigned = errors.New("user pack is not defined")

	// ErrGroupNotFound is returned when a user was not found.
	ErrGroupNotFound = errors.New("group not found")

	// ErrUserNotFound is returned when a user was not found.
	ErrUserNotFound = errors.New("user not found")

	// ErrTokenNotFound is returned when a token was not found.
	ErrTokenNotFound = errors.New("token not found")
)
