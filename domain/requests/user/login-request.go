package user_requests

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
	// TODO: Will update logic
	//pattern := `^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[\W_])^`

	// Compile the regular expression
	//compiledRegexp, _ := regexp.Compile(pattern)

	return validation.ValidateStruct(&ba,
		validation.Field(&ba.Email, is.Email),
		validation.Field(&ba.Password, validation.Length(minPathLength, 0)),
		//validation.Field(&ba.Password, validation.Match(compiledRegexp)),
	)
}

type LoginRequest struct {
	*BasicAuth
}
