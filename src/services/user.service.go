package services

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/requests"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/responses"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
	"github.com/pkg/errors"
)

type IUserService interface {
	SignUp(requests.SignUpRequest) error
	GetUser() (responses.GetUserResponse, error)
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService {
	return &UserService{userRepository}
}

/*
サインアップ
*/
func (userService *UserService) SignUp(req requests.SignUpRequest) error {
	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
	err := userService.userRepository.CreateUser(user)
	if err != nil {
		return errors.Errorf("Something went wrong")
	}
	return nil
}

/*
ユーザ取得
*/
func (userService *UserService) GetUser() (responses.GetUserResponse, error) {
	user, err := userService.userRepository.GetUser()
	if err != nil {
		return responses.GetUserResponse{}, errors.Errorf("Something went wrong")
	}
	return responses.NewGetUserResponse(user), err
}
