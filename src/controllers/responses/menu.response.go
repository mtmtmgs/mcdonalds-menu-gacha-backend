package responses

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/config"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
)

/*
メニューリスト取得
*/
type GetMenuListResponse struct {
	PagingResponse
	Items []GetMenuListItem `json:"items"`
}

type GetMenuListItem struct {
	Id           uint   `json:"id"`
	CreatedAt    string `json:"createdAt"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Category     string `json:"category"`
	MealTimeType string `json:"mealTimeType"`
}

func NewGetMenuListResponse(menuList []models.Menu, totalCount int) GetMenuListResponse {
	res := GetMenuListResponse{
		Items: []GetMenuListItem{},
	}

	for _, menu := range menuList {
		res.TotalCount = totalCount
		res.PerPageCount = config.PerPageCount
		res.Items = append(res.Items, GetMenuListItem{
			Id:           menu.Id,
			CreatedAt:    menu.CreatedAt.Format("2006-01-02 15:04:05"),
			Name:         menu.Name,
			Price:        menu.Price,
			Category:     menu.Category,
			MealTimeType: menu.MealTimeType,
		})
	}
	return res
}

/*
メニューガチャ取得
*/
type GetMenuGachaResponse struct {
	Budget int                `json:"budget"`
	Items  []GetMenuGachaItem `json:"items"`
}

type GetMenuGachaItem struct {
	Id           uint   `json:"id"`
	CreatedAt    string `json:"createdAt"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Category     string `json:"category"`
	MealTimeType string `json:"mealTimeType"`
}

func NewGetMenuGachaResponse(menuList []models.Menu, budget int) GetMenuGachaResponse {
	res := GetMenuGachaResponse{
		Items: []GetMenuGachaItem{},
	}

	for _, menu := range menuList {
		res.Budget = budget
		res.Items = append(res.Items, GetMenuGachaItem{
			Id:           menu.Id,
			CreatedAt:    menu.CreatedAt.Format("2006-01-02 15:04:05"),
			Name:         menu.Name,
			Price:        menu.Price,
			Category:     menu.Category,
			MealTimeType: menu.MealTimeType,
		})
	}
	return res
}
