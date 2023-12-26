package controllers

import (
	"net/http"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/requests"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/services"
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
	return &UserController{userService}
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
	return c.JSON(http.StatusOK, nil)
}

/*
ログイン
*/
func (userController *UserController) Login(c echo.Context) error {
	_, err := requests.NewLoginRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, nil)
}

/*
ユーザ取得
*/
func (userController *UserController) GetUser(c echo.Context) error {
	res, err := userController.userService.GetUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, res)
}
