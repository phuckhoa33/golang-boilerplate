package user_requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type ResetPasswordRequest struct {
	NewPassword        string `json:"newPassword" validate:"required" example:"332003"`
	ConfirmNewPassword string `json:"newConfirmPassword" validate:"required" example:"332003"`
}

func (rpr ResetPasswordRequest) Validate() error {
	return validation.ValidateStruct(&rpr,
		validation.Field(&rpr.NewPassword, validation.Length(minPathLength, 0)),
		validation.Field(&rpr.ConfirmNewPassword, validation.Length(minPathLength, 0)),
	)
}
