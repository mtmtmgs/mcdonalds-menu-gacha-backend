package responses

import "github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/models"

/*
ユーザ取得
*/
type GetUserResponse struct {
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
}

func NewGetUserResponse(user models.User) GetUserResponse {
	res := GetUserResponse{}

	res.LastName = user.LastName
	res.FirstName = user.FirstName
	res.Email = user.Email
	return res
}
