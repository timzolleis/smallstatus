package validations

import "github.com/go-playground/validator/v10"

func ValidateHttpMethod(input validator.FieldLevel) bool {
	validMethods := []string{"GET", "POST", "PUT", "DELETE"}
	for _, method := range validMethods {
		if input.Field().String() == method {
			return true
		}
	}
	return false
}
