package user_requests

type UpdateUserInfoRequest struct {
	Email       string `json:"email" example:"phuckhoa81@gmail.com"`
	Username    string `json:"username" example:"phuckhoa"`
	PhoneNumber string `json:"phoneNumber" example:"842495038"`
	FullName    string `json:"fullName" example:"Nguyen Khoa Minh Phuc"`
	Address     string `json:"address" example:"8 Ward, Binh Chanh District Ho Chi Minh City"`
	Gender      string `json:"gender" example:"MALE"`
	DateOfBirth string `json:"dateOfBirth" example:"03/03/2003"`
}
