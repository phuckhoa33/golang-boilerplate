package user_requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required" example:"phuckhoa81@gmail.com"`
}

func (fpr ForgotPasswordRequest) Validate() error {
	return validation.ValidateStruct(&fpr,
		validation.Field(&fpr.Email, is.Email.Error("EMAIL_IS_INVALID"), validation.Required.Error("EMAIL_IS_REQUIRED")))
}
