package services

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
)

type IUserService interface {
	IsUserEmailDuplication(string) bool
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(
	userRepository repositories.IUserRepository,
) *UserService {
	userService := UserService{
		userRepository: userRepository,
	}
	utils.CheckDependencies(userService)
	return &userService
}

/*
ユーザのメールアドレスが重複しているかどうか
*/
func (userService *UserService) IsUserEmailDuplication(email string) bool {
	_, err := userService.userRepository.GetUserByEmail(email)
	return err == nil
}
