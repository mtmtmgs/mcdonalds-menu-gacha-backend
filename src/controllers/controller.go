package controllers

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/usecases"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
)

type Manager struct {
	UserController IUserController
	MenuController IMenuController
}

func New(usecaseManager *usecases.Manager) *Manager {
	manager := Manager{
		UserController: NewUserController(usecaseManager.UserUsecase),
		MenuController: NewMenuController(usecaseManager.MenuUsecase),
	}
	utils.CheckDependencies(manager)
	return &manager
}
