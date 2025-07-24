package model

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// ##### NOT IMPLEMENTED #####
// Helper code to validate the format of the incoming data
func ValidateTenantInput(input Tenant) error {
	val := validator.New(validator.WithRequiredStructEnabled())
	err := val.Struct(input)

	var errValidation string

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errValidation = e.Field() + " " + e.Tag() + ";" + errValidation
		}
		return errors.New(errValidation)
	}
	return nil
}
