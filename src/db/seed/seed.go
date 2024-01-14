package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/db"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/env"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
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
		menu, err := models.NewMenu(models.Menu{
			Name:         record[0],
			Price:        int64(price),
			Category:     record[2],
			MealTimeType: record[3],
		})
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
