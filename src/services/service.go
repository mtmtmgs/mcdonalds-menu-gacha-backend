package services

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
)

type Manager struct {
	UserService IUserService
	MenuService IMenuService
}

func New(repositoryManager *repositories.Manager) *Manager {
	return &Manager{
		UserService: NewUserService(
			repositoryManager.BaseRepository,
			repositoryManager.UserRepository,
		),
		MenuService: NewMenuService(
			repositoryManager.BaseRepository,
			repositoryManager.MenuRepository,
		),
	}
}
