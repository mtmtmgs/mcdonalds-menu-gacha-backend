package responses

/*
トークン
*/
type TokenResponse struct {
	Token string `json:"token"`
}

func NewTokenResponse(token string) TokenResponse {
	res := TokenResponse{}

	res.Token = token
	return res
}

/*
ページング
*/
type PagingResponse struct {
	TotalCount   int `json:"totalCount"`
	PerPageCount int `json:"perPageCount"`
}
