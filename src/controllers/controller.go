package controllers

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/services"
)

type Manager struct {
	UserController IUserController
	MenuController IMenuController
}

func New(serviceManager *services.Manager) *Manager {
	return &Manager{
		UserController: NewUserController(serviceManager.UserService),
		MenuController: NewMenuController(serviceManager.MenuService),
	}
}
