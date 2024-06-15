package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	minPathLength = 8
)

type BasicAuth struct {
	Email    string `json:"email" validate:"required" example:"phuckhoa81@gmail.com"`
	Password string `json:"password" validate:"required" example:"33003"`
}

func (basicAuth BasicAuth) Validate() error {
	return validation.ValidateStruct(&basicAuth,
		validation.Field(&basicAuth.Email, is.Email),
		validation.Field(&basicAuth.Password, validation.Length(minPathLength, 0)))
}

type LoginRequest struct {
	BasicAuth
}

type RegisterRequest struct {
	BasicAuth
	Name string `json:"name" validate:"required" example:"phuckhoa"`
}

func (registerRequest RegisterRequest) Validate() error {
	err := registerRequest.BasicAuth.Validate()
	if err != nil {
		return err
	}

	return validation.ValidateStruct(&registerRequest, validation.Field(&registerRequest.Name, validation.Required))
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required" example:"refreshToken"`
}
