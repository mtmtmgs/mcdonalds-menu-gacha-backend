package requests

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

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
