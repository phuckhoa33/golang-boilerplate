package user_requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" example:"123456@Abc"`
	NewPassword string `json:"newPassword" example:"$0972495038Phuc"`
}

func (cpq ChangePasswordRequest) Validate() error {

	return validation.ValidateStruct(&cpq,
		validation.Field(
			&cpq.NewPassword,
			validation.Length(minPathLength, 0).Error("PASSWORD_IS_TOO_SHORT"),
			validation.Match(regexp.MustCompile("[A-Z]")).Error("PASSWORD_IS_INVALID"),
			validation.Match(regexp.MustCompile("[a-z]")).Error("PASSWORD_IS_INVALID"),
			validation.Match(regexp.MustCompile("\\d")).Error("PASSWORD_IS_INVALID"),
			validation.Match(regexp.MustCompile(`[\W_]`)).Error("PASSWORD_IS_INVALID"),
		),
	)
}
