package user_responses

import (
	"golang-boilerplate/models"
	"time"
)

type ViewUserProfileResponse struct {
	ID               string    `json:"id" example:"d103f471-0ccd-4858-b417-d67b09910d34"`
	Email            string    `json:"email" example:"phuckhoa81@gmail.com"`
	Username         string    `json:"username" example:"phuckhoa"`
	PhoneNumber      string    `json:"phoneNumber" example:"84972495038"`
	FullName         string    `json:"fullName" example:"Nguyen Khoa Minh Phuc"`
	Address          string    `json:"address" example:"Ho Chi Minh City"`
	Gender           string    `json:"gender" example:"MALE"`
	DateOfBirth      time.Time `json:"dateOfBirth" example:"03/03/03"`
	Avatar           string    `json:"avatar" example:"image.png"`
	RoleId           string    `json:"roleId" example:"d103f471-0ccd-4858-b417-d67b09910d34"`
	VerifyAccountOtp string    `json:"verifyAccountOtp" example:"00100"`
}

func NewViewUserProfileResponse(user *models.User) *ViewUserProfileResponse {
	return &ViewUserProfileResponse{
		ID:               user.ID.String(),
		Email:            user.Email,
		Username:         user.Username,
		PhoneNumber:      user.PhoneNumber,
		FullName:         user.FullName,
		Address:          user.Address,
		Gender:           user.Gender,
		DateOfBirth:      user.DateOfBirth,
		Avatar:           user.Avatar,
		RoleId:           user.RoleId.String(),
		VerifyAccountOtp: user.VerifyAccountOtp,
	}
}
