package controllers

import (
	"net/http"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/services"
	"github.com/labstack/echo/v4"
)

type Menu interface {
	GetMenus(c echo.Context) error
}

type MenuController struct {
	menuService services.Menu
}

func NewMenuController(menuService services.Menu) *MenuController {
	return &MenuController{menuService}
}

// TODO response型
func (menuController *MenuController) GetMenus(c echo.Context) error {
	menus, err := menuController.menuService.GetMenus()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cannot get menus")
	}
	return c.JSONPretty(http.StatusOK, menus, "  ")
}
