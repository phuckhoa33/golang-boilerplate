package user_requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type RefreshRequest struct {
	Token string `json:"token" validate:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJkMTAzZjQ3MS0wY2NkLTQ4NTgtYjQxNy1kNjdiMDk5MTBkMzQiLCJ1c2VybmFtZSI6IkpvaG5Eb2UiLCJleHAiOjE3MjAyNjUxNjd9.MBl3tPb9T-r7QsQTrTHENYd-UvSCzLMN7oKgOEoHxIo"`
}

func (rr RefreshRequest) Validate() error {
	return validation.ValidateStruct(&rr, validation.Field(&rr.Token, validation.Required.Error("TOKEN_IS_REQUIRED")))
}
