package responses

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Expired      int64  `json:"expired"`
}

func NewLoginResponse(token, refreshToken string, expired int64) *LoginResponse {
	return &LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		Expired:      expired,
	}
}
