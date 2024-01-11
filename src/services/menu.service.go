package services

import (
	"math/rand"
	"time"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/requests"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers/responses"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
	"github.com/pkg/errors"
)

type IMenuService interface {
	GetMenuList(requests.GetMenuListRequest) (responses.GetMenuListResponse, error)
	GetMenuGacha(requests.GetMenuGachaRequest) (responses.GetMenuGachaResponse, error)
}

type MenuService struct {
	baseRepository repositories.IBaseRepository
	menuRepository repositories.IMenuRepository
}

func NewMenuService(
	baseRepository repositories.IBaseRepository,
	menuRepository repositories.IMenuRepository,
) *MenuService {
	menuService := MenuService{
		baseRepository: baseRepository,
		menuRepository: menuRepository,
	}
	utils.CheckDependencies(menuService)
	return &menuService
}

/*
メニューリスト取得
*/
func (menuService *MenuService) GetMenuList(req requests.GetMenuListRequest) (responses.GetMenuListResponse, error) {
	var res responses.GetMenuListResponse
	menuList, totalCount, err := menuService.menuRepository.GetMenuList(req.Page, req.Category, req.MealTimeType)
	if err != nil {
		return res, errors.Errorf("Something went wrong")
	}
	res = responses.NewGetMenuListResponse(menuList, totalCount)
	return res, err
}

/*
メニューガチャ取得
*/
func (menuService *MenuService) GetMenuGacha(req requests.GetMenuGachaRequest) (responses.GetMenuGachaResponse, error) {
	var res responses.GetMenuGachaResponse
	menuList, err := repositories.GetList[models.Menu](menuService.baseRepository.GetDB())
	if err != nil {
		return res, errors.Errorf("Something went wrong")
	}

	// 予算内のメニューに絞る
	var menuWithinBudget []models.Menu
	for _, menu := range menuList {
		if int(menu.Price) <= req.Budget {
			menuWithinBudget = append(menuWithinBudget, menu)
		}
	}

	// 乱数シード
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 予算内でメニューを繰り返して格納
	var gachaMenuList []models.Menu
	budget := req.Budget
	for {
		if len(menuWithinBudget) == 0 {
			break
		}
		var tmp []models.Menu
		ran := r.Intn(len(menuWithinBudget))
		budget = budget - int(menuWithinBudget[ran].Price)
		gachaMenuList = append(gachaMenuList, menuWithinBudget[ran])
		for _, menu := range menuWithinBudget {
			if int(menu.Price) <= budget {
				tmp = append(tmp, menu)
			}
		}
		menuWithinBudget = tmp
	}
	res = responses.NewGetMenuGachaResponse(gachaMenuList, req.Budget)
	return res, err
}
