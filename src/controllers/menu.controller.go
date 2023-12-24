package controllers

import (
	"net/http"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/requests"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/services"
	"github.com/labstack/echo/v4"
)

func NewMenuController(menuService services.IMenuService) *MenuController {
	return &MenuController{menuService}
}

type IMenuController interface {
	GetMenuList(c echo.Context) error
}

type MenuController struct {
	menuService services.IMenuService
}

/*
メニューリスト取得
*/
func (menuController *MenuController) GetMenuList(c echo.Context) error {
	req, err := requests.NewGetMenuListRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res, err := menuController.menuService.GetMenuList(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, res)
}
