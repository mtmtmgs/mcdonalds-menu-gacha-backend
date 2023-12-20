package services

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
)

type IMenuService interface {
	GetMenus() ([]models.Menu, error)
}

type MenuService struct {
	menuRepository repositories.IMenuRepository
}

func NewMenuService(menuRepository repositories.IMenuRepository) *MenuService {
	return &MenuService{menuRepository}
}

func (menuService *MenuService) GetMenus() ([]models.Menu, error) {
	menus, err := menuService.menuRepository.GetMenus()
	return menus, err
}
