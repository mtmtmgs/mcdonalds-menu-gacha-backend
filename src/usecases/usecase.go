package usecases

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/services"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
)

type Manager struct {
	UserUsecase IUserUsecase
	MenuUsecase IMenuUsecase
}

func New(
	repositoryManager *repositories.Manager,
	serviceManager *services.Manager,
) *Manager {
	manager := Manager{
		UserUsecase: NewUserUsecase(
			repositoryManager.BaseRepository,
			repositoryManager.UserRepository,
			serviceManager.UserService,
		),
		MenuUsecase: NewMenuUsecase(
			repositoryManager.BaseRepository,
			repositoryManager.MenuRepository,
		),
	}
	utils.CheckDependencies(manager)
	return &manager
}
