package repositories

import "github.com/uptrace/bun"

type Manager struct {
	UserRepository IUserRepository
	MenuRepository IMenuRepository
}

func New(db *bun.DB) *Manager {
	return &Manager{
		UserRepository: NewUserRepository(db),
		MenuRepository: NewMenuRepository(db),
	}
}
