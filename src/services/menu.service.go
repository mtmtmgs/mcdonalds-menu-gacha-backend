package services

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/requests"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/responses"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
	"github.com/pkg/errors"
)

type IMenuService interface {
	GetMenuList(requests.GetMenuListRequest) (responses.GetMenuListResponse, error)
}

type MenuService struct {
	menuRepository repositories.IMenuRepository
}

func NewMenuService(menuRepository repositories.IMenuRepository) *MenuService {
	return &MenuService{menuRepository: menuRepository}
}

/*
メニューリスト取得
*/
func (menuService *MenuService) GetMenuList(req requests.GetMenuListRequest) (responses.GetMenuListResponse, error) {
	var res responses.GetMenuListResponse
	menuList, err := menuService.menuRepository.GetMenuList()
	if err != nil {
		return res, errors.Errorf("Something went wrong")
	}
	res = responses.NewGetMenuListResponse(menuList)
	return res, err
}
