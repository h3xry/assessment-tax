package utils

import (
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type ValidatorError struct {
	Message string
	Fields  map[string]string
}

func (e ValidatorError) Error() string {
	return fmt.Sprint(e.Message)
}

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fl reflect.StructField) string {
		name := strings.SplitN(fl.Tag.Get("json"), ",", 2)
		if len(name) > 1 && name[1] == "-" {
			return ""
		}
		return name[0]
	})
	return validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return ValidatorError{
			Message: "invalid request",
			Fields:  validatorErrors(err),
		}
	}
	return nil
}

func validatorErrors(err error) map[string]string {
	errFields := map[string]string{}
	for _, err := range err.(validator.ValidationErrors) {
		errFields[err.Field()] = fmt.Sprintf(
			"failed '%s' tag check value '%v' is not valid",
			err.Tag(), err.Value(),
		)
	}
	return errFields
}
