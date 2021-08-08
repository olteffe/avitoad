package validators

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// AdValidator func for create a new validator for expected fields,
// register function to get tag name from `json` tags.
func AdValidator() *validator.Validate {
	// Create a new validator.
	v := validator.New()

	// Get tag name from `json`.
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// Define name of field.
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		// Processing name value.
		if name == "-" {
			return ""
		}
		return name
	})

	// Validator for ad name.
	_ = v.RegisterValidation("name", func(fl validator.FieldLevel) bool {
		// Define field as string.
		field := fl.Field().String()

		// Return true, if string length <= 200.
		return len(field) <= 200
	})

	// Validator for ad about.
	_ = v.RegisterValidation("about", func(fl validator.FieldLevel) bool {
		// Define field as string.
		field := fl.Field().String()

		// Return true, if string length <= 1000.
		return len(field) <= 1000
	})

	// Validator for ad photos.
	_ = v.RegisterValidation("photos", func(fl validator.FieldLevel) bool {
		// Define field as stringArray.
		field := fl.Field().Bytes()

		// Return true, if length <= 3.
		return len(field) <= 3
	})

	return v
}
