package repositories

import "github.com/uptrace/bun"

type Manager struct {
	MenuRepository Menu
}

func New(db *bun.DB) *Manager {
	return &Manager{
		MenuRepository: NewMenuRepository(db),
	}
}
