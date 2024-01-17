package entities

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/models"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/values"
)

/*
新規インスタンス生成
新規追加の際はこれを使う
*/
func NewMenu(
	name values.MenuName,
	price values.MenuPrice,
	category values.MenuCategory,
	mealTimeType values.MenuMealTimeType,
) (models.Menu, error) {
	var menu models.Menu
	err := name.Validate()
	if err != nil {
		return menu, err
	}
	err = price.Validate()
	if err != nil {
		return menu, err
	}
	err = category.Validate()
	if err != nil {
		return menu, err
	}
	err = mealTimeType.Validate()
	if err != nil {
		return menu, err
	}

	menu.Name = name.Value()
	menu.Price = price.Value()
	menu.Category = category.Value()
	menu.MealTimeType = mealTimeType.Value()
	return menu, nil
}
