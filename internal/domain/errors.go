// package domain defines the domain models and interfaces for the microservice.
package domain

import "errors"

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrInvalidLimit  = errors.New("invalid limit")
	ErrInvalidOffset = errors.New("invalid offset")
)
