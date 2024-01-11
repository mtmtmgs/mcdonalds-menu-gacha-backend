package services

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
)

type Manager struct {
	UserService IUserService
	MenuService IMenuService
}

func New(repositoryManager *repositories.Manager) *Manager {
	manager := Manager{
		UserService: NewUserService(
			repositoryManager.BaseRepository,
			repositoryManager.UserRepository,
		),
		MenuService: NewMenuService(
			repositoryManager.BaseRepository,
			repositoryManager.MenuRepository,
		),
	}
	utils.CheckDependencies(manager)
	return &manager
}
