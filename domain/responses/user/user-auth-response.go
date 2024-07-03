package user_responses

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Exp          int64  `json:"exp"`
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
