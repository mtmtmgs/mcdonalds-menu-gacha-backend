package repositories

import (
	"context"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/consts"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
	"github.com/uptrace/bun"
)

type IMenuRepository interface {
	GetMenuList(int, string, string) ([]models.Menu, int, error)
}

type MenuRepository struct {
	db *bun.DB
}

func NewMenuRepository(db *bun.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (menuRepository *MenuRepository) GetMenuList(page int, category string, mealTimeType string) ([]models.Menu, int, error) {
	var menuList []models.Menu
	var totalCount int

	ctx := context.Background()
	query := menuRepository.db.NewSelect().Model(&menuList)
	totalCount, err := query.Count(ctx)
	if err != nil {
		return menuList, totalCount, err
	}
	if page != 0 {
		query.Offset((page - 1) * consts.PerPageCount)
	}
	if category != "" {
		query.Where("category = ?", category)
	}
	if mealTimeType != "" {
		query.Where("meal_time_type = ?", mealTimeType)
	}
	err = query.Limit(consts.PerPageCount).Scan(ctx)
	return menuList, totalCount, err
}
