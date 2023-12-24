package repositories

import (
	"context"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
	"github.com/uptrace/bun"
)

type IMenuRepository interface {
	GetMenuList() ([]models.Menu, error)
}

type MenuRepository struct {
	db *bun.DB
}

func NewMenuRepository(db *bun.DB) *MenuRepository {
	return &MenuRepository{db}
}

func (menuRepository MenuRepository) GetMenuList() ([]models.Menu, error) {
	var menuList []models.Menu
	ctx := context.Background()
	err := menuRepository.db.NewSelect().Model(&menuList).Scan(ctx)
	return menuList, err
}
