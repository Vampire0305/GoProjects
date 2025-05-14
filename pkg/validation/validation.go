package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) error {
	if errs, ok := err.(validator.ValidationErrors); ok {
		msg := "validation failed:"
		for _, fieldErr := range errs {
			msg += fmt.Sprintf(" %s (%s);", fieldErr.Field(), fieldErr.Tag())
		}
		return fmt.Errorf(msg)
	}
	return err
}
