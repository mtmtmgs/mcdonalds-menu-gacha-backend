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
	return &MenuService{menuRepository}
}

/*
メニューリスト取得
*/
func (menuService *MenuService) GetMenuList(req requests.GetMenuListRequest) (responses.GetMenuListResponse, error) {
	menuList, err := menuService.menuRepository.GetMenuList()
	if err != nil {
		return responses.GetMenuListResponse{}, errors.Errorf("Something went wrong")
	}
	return responses.NewGetMenuListResponse(menuList), err
}
