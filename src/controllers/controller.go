package controllers

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/services"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
)

type Manager struct {
	UserController IUserController
	MenuController IMenuController
}

func New(serviceManager *services.Manager) *Manager {
	manager := Manager{
		UserController: NewUserController(serviceManager.UserService),
		MenuController: NewMenuController(serviceManager.MenuService),
	}
	utils.CheckDependencies(manager)
	return &manager
}
