package user_requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang-boilerplate/domain/enums"
	"regexp"
)

type UpdateUserInfoRequest struct {
	Email       string `json:"email" example:"phuckhoa81@gmail.com"`
	Username    string `json:"username" example:"phuckhoa"`
	PhoneNumber string `json:"phoneNumber" example:"842495038"`
	FullName    string `json:"fullName" example:"Nguyen Khoa Minh Phuc"`
	Address     string `json:"address" example:"8 Ward, Binh Chanh District Ho Chi Minh City"`
	Gender      string `json:"gender" example:"MALE"`
	DateOfBirth string `json:"dateOfBirth" example:"03/03/2003"`
}

func (r UpdateUserInfoRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required.Error("EMAIL_IS_REQUIRED"), is.Email.Error("EMAIL_IS_INVALID")),
		validation.Field(&r.Username, validation.Required.Error("USERNAME_IS_REQUIRED")),
		validation.Field(&r.PhoneNumber, validation.Required.Error("PHONE_NUMBER_IS_REQUIRED"), validation.Match(regexp.MustCompile(`^\+[1-9]\d{1,14}$`)).Error("PHONE_NUMBER_IS_INVALID")),
		validation.Field(&r.FullName, validation.Required.Error("FULL_NAME_IS_REQUIRED")),
		validation.Field(&r.DateOfBirth, validation.Required.Error("DATE_OF_BIRTH_IS_REQUIRED")),
		validation.Field(&r.Address, validation.Required.Error("ADDRESS_IS_REQUIRED")),
		validation.Field(&r.Gender, validation.Required.Error("GENDER_IS_REQUIRED"), validation.In(enums.FEMALE, enums.MALE, enums.OTHERS)))
}
