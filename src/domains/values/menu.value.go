package values

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/consts"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
	"github.com/pkg/errors"
)

/*
メニュー名
*/
type MenuName struct {
	value string
}

func NewMenuName(value string) MenuName {
	return MenuName{value: value}
}

func (mn *MenuName) Validate() error {
	return nil
}

func (mn *MenuName) Value() string {
	return mn.value
}

/*
メニュー価格
*/
type MenuPrice struct {
	value int64
}

func NewMenuPrice(value int64) MenuPrice {
	return MenuPrice{value: value}
}

func (mp *MenuPrice) Validate() error {
	return nil
}

func (mp *MenuPrice) Value() int64 {
	return mp.value
}

/*
メニューカテゴリ
*/
type MenuCategory struct {
	value string
}

func NewMenuCategory(value string) MenuCategory {
	return MenuCategory{value: value}
}

func (mc *MenuCategory) Validate() error {
	categories := []string{consts.Burger, consts.Set, consts.Side, consts.Drink, consts.HappySet, consts.Dessert, consts.McCafe}
	if !utils.IsSliceContains[string](categories, mc.value) {
		return errors.Errorf("カテゴリが不正です")
	}
	return nil
}

func (mc *MenuCategory) Value() string {
	return mc.value
}

/*
メニュー時間帯
*/
type MenuMealTimeType struct {
	value string
}

func NewMenuMealTimeType(value string) MenuMealTimeType {
	return MenuMealTimeType{value: value}
}

func (mmty *MenuMealTimeType) Validate() error {
	mealTimeTypes := []string{consts.Morning, consts.Noon, consts.Night, consts.Regular}
	if !utils.IsSliceContains[string](mealTimeTypes, mmty.value) {
		return errors.Errorf("時間帯が不正です")
	}
	return nil
}

func (mmty *MenuMealTimeType) Value() string {
	return mmty.value
}
