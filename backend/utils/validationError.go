package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationError(req any, err error) map[string]string {
	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		tag := e.Tag()

		t := reflect.TypeOf(req)

		if structField, ok := t.FieldByName(field); ok {
			jsonTag := structField.Tag.Get("json")

			if jsonTag != "" {
				field = jsonTag
			} else {
				field = strings.ToLower(field)
			}
		}

		switch tag {
		case "required":
			errors[field] = "field is required"
		case "email":
			errors[field] = "invalid email format"
		case "min":
			errors[field] = "minimum " + e.Param() + " characters"
		case "phone":
			errors[field] = "invalid phone number"
		default:
			errors[field] = "invalid value"
		}
	}
	return errors
}
