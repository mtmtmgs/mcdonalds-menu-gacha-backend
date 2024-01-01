package requests

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

/*
ページング
*/
type PagingRequest struct {
	Page int
}

func NewPagingRequest(c echo.Context) PagingRequest {
	req := PagingRequest{}

	pageNum, err := strconv.Atoi(c.QueryParam("page"))
	if err == nil {
		req.Page = pageNum
	}
	return req
}
