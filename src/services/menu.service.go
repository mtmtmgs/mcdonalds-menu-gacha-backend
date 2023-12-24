package services

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/requests"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/responses"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
)

type IMenuService interface {
	GetMenuList(requests.GetMenuListRequest) (responses.GetMenuListResponse, error)
}

type MenuService struct {
	menuRepository repositories.IMenuRepository
}

func NewMenuService(menuRepository repositories.IMenuRepository) *MenuService {
	return &MenuService{menuRepository}
}

/*
メニューリスト取得
*/
func (menuService *MenuService) GetMenuList(req requests.GetMenuListRequest) (responses.GetMenuListResponse, error) {
	menuList, err := menuService.menuRepository.GetMenuList()
	return responses.NewGetMenuListResponse(menuList), err
}
