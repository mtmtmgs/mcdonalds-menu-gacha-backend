package router

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}

func Register(r *echo.Echo, controllerManager *controllers.Manager) {
	rv := r.Group("/v1")
	rv.GET("/menus", controllerManager.MenuController.GetMenuList)
}
