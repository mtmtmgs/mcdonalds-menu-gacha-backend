package responses

type TokenResponse struct {
	Token string `json:"token"`
}

func NewTokenResponse(token string) TokenResponse {
	res := TokenResponse{}

	res.Token = token
	return res
}
