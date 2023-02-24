package requests

import (
	"errors"
	"fish-hunter/businesses/samples"

	"github.com/go-playground/validator/v10"
)

type UpdateSamples struct {
	Brands []string `json:"brands" validate:"required"`
	Language string `json:"language" validate:"required"`
	Details string `json:"details" validate:"required"`
	Type string `json:"type" validate:"required"`
}

func (u *UpdateSamples) ToDomain() *samples.Domain {
	return &samples.Domain{
		Brands: u.Brands,
		Language: u.Language,
		Details: u.Details,
		Type: u.Type,
	}
}

func (u *UpdateSamples) Validate() error {
	validate := validator.New()
	if validate.Struct(u) != nil {
		return errors.New("some of the fields are not valid")
	}
	return nil
}