package services

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
)

type Manager struct {
	MenuService Menu
}

func New(repositoryManager *repositories.Manager) *Manager {
	return &Manager{
		MenuService: NewMenuService(repositoryManager.MenuRepository),
	}
}
