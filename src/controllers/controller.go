package controllers

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/services"
)

type Manager struct {
	MenuController IMenuController
}

func New(serviceManager *services.Manager) *Manager {
	return &Manager{
		MenuController: NewMenuController(serviceManager.MenuService),
	}
}
