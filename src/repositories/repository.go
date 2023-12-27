package repositories

import "github.com/uptrace/bun"

type Manager struct {
	BaseRepository IBaseRepository
	UserRepository IUserRepository
	MenuRepository IMenuRepository
}

func New(db *bun.DB) *Manager {
	return &Manager{
		BaseRepository: NewBaseRepository(db),
		UserRepository: NewUserRepository(db),
		MenuRepository: NewMenuRepository(db),
	}
}
