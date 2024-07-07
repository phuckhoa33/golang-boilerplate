package user_requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" example:"123456@Abc"`
	NewPassword string `json:"newPassword" example:"$0972495038Phuc"`
}

func (cpq ChangePasswordRequest) Validate() error {
	// TODO: Will update logic
	//pattern := `^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[\W_])^`

	// Compile the regular expression
	//compiledRegexp, _ := regexp.Compile(pattern)

	return validation.ValidateStruct(&cpq,
		validation.Field(&cpq.NewPassword, validation.Length(minPathLength, 0)),
		//validation.Field(&cpq.NewPassword, validation.Match(compiledRegexp)),
	)
}
