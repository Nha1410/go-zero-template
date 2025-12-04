package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Validate validates a struct
func Validate(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		return formatValidationError(err)
	}
	return nil
}

// formatValidationError formats validation errors into a readable message
func formatValidationError(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var messages []string
		for _, fieldError := range validationErrors {
			message := fmt.Sprintf("Field '%s' failed validation: %s", fieldError.Field(), getValidationMessage(fieldError))
			messages = append(messages, message)
		}
		return fmt.Errorf("validation failed: %s", strings.Join(messages, "; "))
	}
	return err
}

// getValidationMessage returns a human-readable validation message
func getValidationMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email address"
	case "min":
		return fmt.Sprintf("must be at least %s", fieldError.Param())
	case "max":
		return fmt.Sprintf("must be at most %s", fieldError.Param())
	case "len":
		return fmt.Sprintf("must be exactly %s characters", fieldError.Param())
	case "oneof":
		return fmt.Sprintf("must be one of: %s", fieldError.Param())
	default:
		return fmt.Sprintf("failed validation rule: %s", fieldError.Tag())
	}
}

