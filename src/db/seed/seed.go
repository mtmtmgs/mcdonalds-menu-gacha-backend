package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/db"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("------------------------------------------------------------")
	log.Println("BEGIN seed")
	log.Println("------------------------------------------------------------")
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	db := db.New()
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

	var menu models.Menu
	var menuList []models.Menu
	for i, record := range records {
		if i == 0 {
			continue
		}
		price, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}
		menu.Name = record[0]
		menu.Price = int64(price)
		menu.Category = record[2]
		menu.MealTimeType = record[3]
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
