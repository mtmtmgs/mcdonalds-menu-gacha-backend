package models

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/consts"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

type Menu struct {
	bun.BaseModel `bun:"table:menus,alias:m"`

	Base
	Name         string `bun:"name,notnull"`
	Price        int64  `bun:"price,nullzero"`
	Category     string `bun:"category"`
	MealTimeType string `bun:"meal_time_type"`
}

func NewMenu(menu Menu) (Menu, error) {
	var newMenu Menu
	name, err := NewMenuName(menu.Name)
	if err != nil {
		return newMenu, err
	}
	price, err := NewMenuPrice(menu.Price)
	if err != nil {
		return newMenu, err
	}
	category, err := NewMenuCategory(menu.Category)
	if err != nil {
		return newMenu, err
	}
	mealTimeType, err := NewMenuMealTimeType(menu.MealTimeType)
	if err != nil {
		return newMenu, err
	}

	newMenu.Name = name
	newMenu.Price = price
	newMenu.Category = category
	newMenu.MealTimeType = mealTimeType
	return newMenu, nil
}

func NewMenuName(name string) (string, error) {
	return name, nil
}

func NewMenuPrice(price int64) (int64, error) {
	return price, nil
}

func NewMenuCategory(category string) (string, error) {
	categories := []string{consts.Burger, consts.Set, consts.Side, consts.Drink, consts.HappySet, consts.Dessert, consts.McCafe}
	if !utils.IsSliceContains[string](categories, category) {
		return "", errors.Errorf("カテゴリが不正です")
	}
	return category, nil
}

func NewMenuMealTimeType(mealTimeType string) (string, error) {
	mealTimeTypes := []string{consts.Morning, consts.Noon, consts.Night, consts.Regular}
	if !utils.IsSliceContains[string](mealTimeTypes, mealTimeType) {
		return "", errors.Errorf("時間帯が不正です")
	}
	return mealTimeType, nil
}
