package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/db"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/entities"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/models"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/values"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/env"
)

func main() {
	log.Println("------------------------------------------------------------")
	log.Println("BEGIN seed")
	log.Println("------------------------------------------------------------")
	env := env.NewCmdEnv("../../.env")
	db := db.New(env)
	defer db.Close()

	file, err := os.Open("./csv/menus.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var menuList []models.Menu
	for i, record := range records {
		if i == 0 {
			continue
		}
		price, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}
		menu, err := entities.NewMenu(
			values.NewMenuName(record[0]),
			values.NewMenuPrice(int64(price)),
			values.NewMenuCategory(record[2]),
			values.NewMenuMealTimeType(record[3]),
		)
		if err != nil {
			log.Printf("追加できませんでした: %v %s", menu, err)
			continue
		}
		menuList = append(menuList, menu)
	}

	ctx := context.Background()
	_, err = db.NewInsert().Model(&menuList).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("------------------------------------------------------------")
	log.Println("FINISHED SUCCESSFUL seed")
	log.Println("------------------------------------------------------------")
}
