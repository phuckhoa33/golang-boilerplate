package user_responses

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
