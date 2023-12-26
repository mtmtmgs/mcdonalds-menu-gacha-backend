package requests

import (
	"github.com/labstack/echo/v4"
)

type SignUpRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func NewSignUpRequest(c echo.Context) (SignUpRequest, error) {
	req := SignUpRequest{}

	err := c.Bind(&req)
	return req, err
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func NewLoginRequest(c echo.Context) (LoginRequest, error) {
	req := LoginRequest{}

	err := c.Bind(&req)
	return req, err
}
