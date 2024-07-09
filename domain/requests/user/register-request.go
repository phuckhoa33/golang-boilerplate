package user_requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang-boilerplate/domain/enums"
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
		validation.Field(&rr.Username, validation.Required.Error("USERNAME_IS_REQUIRED")),
		validation.Field(&rr.Gender, validation.Required.Error("GENDER_IS_REQUIRED"), validation.In(enums.FEMALE, enums.MALE, enums.OTHERS)),
		validation.Field(&rr.FullName, validation.Required.Error("FULL_NAME_IS_REQUIRED")),
	)
}
