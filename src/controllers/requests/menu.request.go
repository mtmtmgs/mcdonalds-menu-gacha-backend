package requests

import (
	"strconv"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/config"
	"github.com/labstack/echo/v4"
)

/*
メニューリスト取得
*/
type GetMenuListRequest struct {
	PagingRequest
	Category     string
	MealTimeType string
}

func NewGetMenuListRequest(c echo.Context) GetMenuListRequest {
	req := GetMenuListRequest{}

	pageNum, err := strconv.Atoi(c.QueryParam("page"))
	if err == nil {
		req.Page = pageNum
	}
	req.Category = c.QueryParam("category")
	req.MealTimeType = c.QueryParam("mealTimeType")
	return req
}

/*
メニューガチャ取得
*/
type GetMenuGachaRequest struct {
	Budget int
}

func NewGetMenuGachaRequest(c echo.Context) GetMenuGachaRequest {
	req := GetMenuGachaRequest{
		Budget: config.MenuGachaDefaultBudget,
	}

	budgetNum, err := strconv.Atoi(c.QueryParam("budget"))
	if err == nil && budgetNum > 0 {
		req.Budget = budgetNum
	}
	return req
}
