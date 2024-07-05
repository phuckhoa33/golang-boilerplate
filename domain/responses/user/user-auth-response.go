package user_responses

type LoginResponse struct {
	AccessToken  string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string `json:"refreshToken" example:"def50200c1c92d..."`
	Exp          int64  `json:"exp" example:"1623180000"`
}

func NewLoginResponse(token, refreshToken string, exp int64) *LoginResponse {
	return &LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		Exp:          exp,
	}
}

type RefreshTokenResponses struct {
	LoginResponse
}

func NewRefreshTokenResponse(token, refreshToken string, exp int64) *RefreshTokenResponses {
	return &RefreshTokenResponses{
		LoginResponse{
			AccessToken:  token,
			RefreshToken: refreshToken,
			Exp:          exp,
		},
	}
}
