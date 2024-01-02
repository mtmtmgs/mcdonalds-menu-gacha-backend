package requests

import (
	"github.com/go-playground/validator/v10"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
	"github.com/labstack/echo/v4"
)

/*
サインアップ
*/
type SignUpRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6,passwordRule"`
}

func NewSignUpRequest(c echo.Context) (SignUpRequest, error) {
	req := SignUpRequest{}

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

/*
ログイン
*/
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func NewLoginRequest(c echo.Context) (LoginRequest, error) {
	req := LoginRequest{}

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
