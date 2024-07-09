package user_requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"regexp"
)

const (
	minPathLength = 8
)

type BasicAuth struct {
	Email    string `json:"email" validate:"required" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required" example:"11111111"`
}

func (ba BasicAuth) Validate() error {
	return validation.ValidateStruct(&ba,
		validation.Field(
			&ba.Email,
			is.Email.Error("EMAIL_IS_INVALID"),
			validation.Required.Error("EMAIL_IS_REQUIRED"),
		),
		validation.Field(
			&ba.Password,
			validation.Length(minPathLength, 0).Error("PASSWORD_IS_TOO_SHORT"),
			validation.Required.Error("PASSWORD_IS_REQUIRED"),
			validation.Match(regexp.MustCompile("[A-Z]")).Error("PASSWORD_IS_INVALID"),
			validation.Match(regexp.MustCompile("[a-z]")).Error("PASSWORD_IS_INVALID"),
			validation.Match(regexp.MustCompile("\\d")).Error("PASSWORD_IS_INVALID"),
			validation.Match(regexp.MustCompile(`[\W_]`)).Error("PASSWORD_IS_INVALID"),
		),
	)
}

type LoginRequest struct {
	*BasicAuth
}
