package repositories

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
	"github.com/uptrace/bun"
)

type Manager struct {
	BaseRepository IBaseRepository
	UserRepository IUserRepository
	MenuRepository IMenuRepository
}

func New(db *bun.DB) *Manager {
	manager := Manager{
		BaseRepository: NewBaseRepository(db),
		UserRepository: NewUserRepository(db),
		MenuRepository: NewMenuRepository(db),
	}
	utils.CheckDependencies(manager)
	return &manager
}
