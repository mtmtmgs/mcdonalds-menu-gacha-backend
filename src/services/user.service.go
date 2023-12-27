package services

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/requests"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/responses"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/router/middlewares"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
	"github.com/pkg/errors"
)

type IUserService interface {
	SignUp(requests.SignUpRequest) error
	Login(requests.LoginRequest) (responses.TokenResponse, error)
	GetUser() (responses.GetUserResponse, error)
}

type UserService struct {
	baseRepository repositories.IBaseRepository
	userRepository repositories.IUserRepository
}

func NewUserService(
	baseRepository repositories.IBaseRepository,
	userRepository repositories.IUserRepository,
) *UserService {
	return &UserService{baseRepository, userRepository}
}

/*
サインアップ
*/
func (userService *UserService) SignUp(req requests.SignUpRequest) error {
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return errors.Errorf("Something went wrong")
	}
	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashPassword,
	}
	err = userService.userRepository.CreateUser(user)
	if err != nil {
		return errors.Errorf("Something went wrong")
	}
	return nil
}

/*
ログイン
*/
func (userService *UserService) Login(req requests.LoginRequest) (responses.TokenResponse, error) {
	user, err := userService.userRepository.GetUserByEmail(req.Email)
	if user == (models.User{}) {
		return responses.TokenResponse{}, errors.Errorf("Unauthorized")
	}
	if err != nil {
		return responses.TokenResponse{}, errors.Errorf("Something went wrong")
	}
	ok := utils.ComparePassword(user.Password, req.Password)
	if !ok {
		return responses.TokenResponse{}, errors.Errorf("Unauthorized")
	}
	// jwt生成
	token, err := middlewares.GenerateJwt(user.Id)
	if err != nil {
		return responses.TokenResponse{}, errors.Errorf("Something went wrong")
	}
	return responses.NewTokenResponse(token), nil
}

/*
ユーザ取得
*/
func (userService *UserService) GetUser() (responses.GetUserResponse, error) {
	// TODO jwtから取得
	user, err := repositories.GetById[models.User](userService.baseRepository.GetDB(), 1)
	if err != nil {
		return responses.GetUserResponse{}, errors.Errorf("Something went wrong")
	}
	return responses.NewGetUserResponse(user), err
}
