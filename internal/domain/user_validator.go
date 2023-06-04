package domain

import "github.com/go-playground/validator/v10"

// userValidator contains the validation logic for the User entity.
type userValidator struct {
	validate *validator.Validate
}

// NewUserValidator creates a new userValidator.
func NewUserValidator() *userValidator {
	validate := validator.New()
	return &userValidator{validate}
}

// ValidateUser validates the given User entity.
func (v *userValidator) ValidateUser(user *User) error {
	return v.validate.Struct(user)
}
