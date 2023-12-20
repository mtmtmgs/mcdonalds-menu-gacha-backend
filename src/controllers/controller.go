package controllers

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/services"
)

type Manager struct {
	MenuController Menu
}

func New(serviceManager *services.Manager) *Manager {
	return &Manager{
		MenuController: NewMenuController(serviceManager.MenuService),
	}
}
