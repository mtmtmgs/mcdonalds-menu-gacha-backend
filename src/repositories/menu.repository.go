package repositories

import (
	"context"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
	"github.com/uptrace/bun"
)

type IMenuRepository interface {
	GetMenus() ([]models.Menu, error)
}

type MenuRepository struct {
	db *bun.DB
}

func NewMenuRepository(db *bun.DB) *MenuRepository {
	return &MenuRepository{db}
}

func (menuRepository MenuRepository) GetMenus() ([]models.Menu, error) {
	var menus []models.Menu
	ctx := context.Background()
	err := menuRepository.db.NewSelect().Model(&menus).Scan(ctx)
	return menus, err
}
