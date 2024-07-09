package user_responses

import (
	"fmt"
	"golang-boilerplate/config"
	"golang-boilerplate/domain/enums"
	"golang-boilerplate/domain/models/postgresql"
	"golang-boilerplate/domain/responses/shared"
	"time"
)

type ViewUserProfileResponse struct {
	ID               string           `json:"id" example:"d103f471-0ccd-4858-b417-d67b09910d34"`
	Email            string           `json:"email" example:"phuckhoa81@gmail.com"`
	Username         string           `json:"username" example:"phuckhoa"`
	PhoneNumber      string           `json:"phoneNumber" example:"84972495038"`
	FullName         string           `json:"fullName" example:"Nguyen Khoa Minh Phuc"`
	Address          string           `json:"address" example:"Ho Chi Minh City"`
	Gender           enums.GenderEnum `json:"gender" example:"MALE"`
	DateOfBirth      time.Time        `json:"dateOfBirth" example:"03/03/03"`
	Avatar           shared.ViewFileResponse
	RoleId           string    `json:"roleId" example:"d103f471-0ccd-4858-b417-d67b09910d34"`
	VerifyAccountOtp string    `json:"verifyAccountOtp" example:"00100"`
	CreatedAt        time.Time `json:"createdAt" example:"2021-07-01T00:00:00Z"`
	UpdatedAt        time.Time `json:"updatedAt" example:"2021-07-01T00:00:00Z"`
	DeleteAt         time.Time `json:"deleteAt" example:"2021-07-01T00:00:00Z"`
}

func NewViewUserProfileResponse(config *config.Config, user *postgresql.User) *ViewUserProfileResponse {
	var avatar shared.ViewFileResponse

	// Check user avatar is existed
	if user.Avatar != "" {
		avatar = shared.ViewFileResponse{
			FileName: user.Avatar,
			FileURL:  fmt.Sprintf("%s/%s/%s", config.Minio.MinioURL, enums.MinioFolderEnumUser, user.Avatar),
		}
	}

	return &ViewUserProfileResponse{
		ID:               user.ID.String(),
		Email:            user.Email,
		Username:         user.Username,
		PhoneNumber:      user.PhoneNumber,
		FullName:         user.FullName,
		Address:          user.Address,
		Gender:           enums.GenderEnum(user.Gender),
		DateOfBirth:      user.DateOfBirth,
		Avatar:           avatar,
		RoleId:           user.RoleId.String(),
		VerifyAccountOtp: user.VerifyAccountOtp,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
		DeleteAt:         user.DeleteAt,
	}
}
