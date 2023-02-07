package validator

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func MapErrorMessage(err error) []string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]string, len(ve))
		for i, fe := range ve {
			out[i] = fe.Field() + ":" + msgForTag(fe)
		}
		return out
	}

	return []string{}
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return fe.Error() // default error
}
