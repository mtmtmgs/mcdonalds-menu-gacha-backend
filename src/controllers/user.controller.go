package controllers

import (
	"net/http"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/requests"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/router/middlewares"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/services"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	GetUser(c echo.Context) error
}

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	userController := UserController{userService: userService}
	utils.CheckDependencies(userController)
	return &userController
}

/*
サインアップ
*/
func (userController *UserController) SignUp(c echo.Context) error {
	req, err := requests.NewSignUpRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = userController.userService.SignUp(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, nil, "  ")
}

/*
ログイン
*/
func (userController *UserController) Login(c echo.Context) error {
	req, err := requests.NewLoginRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res, err := userController.userService.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	return c.JSONPretty(http.StatusOK, res, "  ")
}

/*
ユーザ取得
*/
func (userController *UserController) GetUser(c echo.Context) error {
	userId := middlewares.GetUserIdByJwt(c)
	res, err := userController.userService.GetUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, nil)
	}
	return c.JSONPretty(http.StatusOK, res, "  ")
}
