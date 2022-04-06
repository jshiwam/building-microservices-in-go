package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", skuValidate)
	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	err := v.validate.Struct(i)

	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)

	var returnErrs []ValidationError
	for _, err := range errs {
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}
	return returnErrs
}

func skuValidate(f1 validator.FieldLevel) bool {
	re := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]+")
	matches := re.FindAllString(f1.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
}
