package user_requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

type ResetPasswordRequest struct {
	NewPassword        string `json:"newPassword" validate:"required" example:"332003"`
	ConfirmNewPassword string `json:"newConfirmPassword" validate:"required" example:"332003"`
}

func (rpr ResetPasswordRequest) Validate() error {
	return validation.ValidateStruct(&rpr,
		validation.Field(
			&rpr.NewPassword,
			validation.Length(minPathLength, 0).Error("PASSWORD_IS_TOO_SHORT"),
			validation.Length(minPathLength, 0).Error("PASSWORD_IS_TOO_SHORT"),
			validation.Match(regexp.MustCompile("[A-Z]")).Error("PASSWORD_IS_INVALID"),
			validation.Match(regexp.MustCompile("[a-z]")).Error("PASSWORD_IS_INVALID"),
			validation.Match(regexp.MustCompile("\\d")).Error("PASSWORD_IS_INVALID"),
			validation.Match(regexp.MustCompile(`[\W_]`)).Error("PASSWORD_IS_INVALID"),
		),
		validation.Field(
			&rpr.ConfirmNewPassword,
			validation.Length(minPathLength, 0).Error("PASSWORD_IS_TOO_SHORT"),
			validation.Match(regexp.MustCompile("[A-Z]")).Error("PASSWORD_IS_INVALID"),
			validation.Match(regexp.MustCompile("[a-z]")).Error("PASSWORD_IS_INVALID"),
			validation.Match(regexp.MustCompile("\\d")).Error("PASSWORD_IS_INVALID"),
			validation.Match(regexp.MustCompile(`[\W_]`)).Error("PASSWORD_IS_INVALID"),
		),
	)
}
