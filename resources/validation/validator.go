package validation

import (
	"github.com/go-playground/validator"
)

type MedisyncValidator struct {
	Validator *validator.Validate
}

func NewMedisyncValidator() *MedisyncValidator {
	return &MedisyncValidator{Validator: validator.New()}
}

func (cv *MedisyncValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
