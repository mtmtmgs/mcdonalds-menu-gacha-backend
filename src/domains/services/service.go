package services

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
)

type Manager struct {
	UserService IUserService
}

func New(repositoryManager *repositories.Manager) *Manager {
	manager := Manager{
		UserService: NewUserService(
			repositoryManager.UserRepository,
		),
	}
	utils.CheckDependencies(manager)
	return &manager
}
