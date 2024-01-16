package controllers

import (
	"net/http"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/requests"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/usecases"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
	"github.com/labstack/echo/v4"
)

type IMenuController interface {
	GetMenuList(c echo.Context) error
	GetMenuGacha(c echo.Context) error
}

type MenuController struct {
	menuUsecase usecases.IMenuUsecase
}

func NewMenuController(menuUsecase usecases.IMenuUsecase) *MenuController {
	menuController := MenuController{menuUsecase: menuUsecase}
	utils.CheckDependencies(menuController)
	return &menuController
}

/*
メニューリスト取得
*/
func (menuController *MenuController) GetMenuList(c echo.Context) error {
	req := requests.NewGetMenuListRequest(c)
	res, err := menuController.menuUsecase.GetMenuList(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, res, "  ")
}

/*
メニューガチャ取得
*/
func (menuController *MenuController) GetMenuGacha(c echo.Context) error {
	req := requests.NewGetMenuGachaRequest(c)
	res, err := menuController.menuUsecase.GetMenuGacha(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, res, "  ")
}
