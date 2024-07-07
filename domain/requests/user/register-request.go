package user_requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RegisterRequest struct {
	*BasicAuth
	Username string `json:"username" validate:"required" example:"JohnDoe"`
	FullName string `json:"fullName" validate:"required" example:"John Doe"`
	Gender   string `json:"gender" validate:"required" example:"MALE"`
}

func (rr RegisterRequest) Validate() error {
	err := rr.BasicAuth.Validate()
	if err != nil {
		return err
	}

	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Username, validation.Required),
	)
}
