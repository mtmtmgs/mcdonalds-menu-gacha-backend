package requests

import (
	"github.com/labstack/echo/v4"
)

type GetMenuListRequest struct {
	Category string
}

func NewGetMenuListRequest(c echo.Context) (GetMenuListRequest, error) {
	req := GetMenuListRequest{}

	req.Category = c.QueryParam("category")
	return req, nil
}
