package user_responses

import "net/url"

type GetUserPutPreSignedPutURLResponse struct {
	PreSignedPutURL string `json:"preSignedPutUrl" example:"https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies f"`
}

func NewGetUserPutPreSignedPutURLResponse(path *url.URL) *GetUserPutPreSignedPutURLResponse {
	return &GetUserPutPreSignedPutURLResponse{
		PreSignedPutURL: path.String(),
	}
}
