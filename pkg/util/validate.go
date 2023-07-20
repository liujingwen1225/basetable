package util

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.SetTagName("valid")
}
func Valid(s interface{}) error {
	return Validate.Struct(s)
}
