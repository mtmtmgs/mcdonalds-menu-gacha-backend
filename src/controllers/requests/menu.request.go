package requests

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/config"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
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
	Budget int `json:"budget"`
}

func NewGetMenuGachaRequest(c echo.Context) (GetMenuGachaRequest, error) {
	req := GetMenuGachaRequest{
		Budget: config.MenuGachaDefaultBudget,
	}

	err := c.Bind(&req)
	if err != nil {
		return req, err
	}
	err = c.Validate(&req)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		err = utils.ConvertValidationErrorMessage(errs)
		return req, err
	}
	return req, err
}
