package service

import (
	"github.com/mholt/binding"
)

func newRequiredField(name string) binding.Field {
	return binding.Field{
		Form:     name,
		Required: true,
	}
}

func newOptionalField(name string) binding.Field {
	return binding.Field{
		Form: name,
	}
}

const (
	bindingValidationError = "ValidationError"
)

func newValidationError(field, message string) binding.Error {
	return binding.Error{
		FieldNames:     []string{field},
		Classification: bindingValidationError,
		Message:        message,
	}
}
