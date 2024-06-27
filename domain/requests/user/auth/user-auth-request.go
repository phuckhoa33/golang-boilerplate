package user_auth_requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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
		validation.Field(&ba.Email, is.Email),
		validation.Field(&ba.Password, validation.Length(minPathLength, 0)),
	)
}

type LoginRequest struct {
	BasicAuth
}

type RegisterRequest struct {
	BasicAuth
	Username string `json:"username" validate:"required" example:"JohnDoe"`
	Fullname string `json:"fullname" validate:"required" example:"John Doe"`
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

type RefreshRequest struct {
	Token string `json:"token" validate:"required" example:"refresh_token"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required" example:"phuckhoa81@gmail.com"`
}

func (fpr ForgotPasswordRequest) Validate() error {
	return validation.ValidateStruct(&fpr,
		validation.Field(&fpr.Email, is.Email),
	)
}
