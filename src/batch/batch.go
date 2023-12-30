package main

import (
	"log"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/batch/scripts"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/db"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("----------batch START----------")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	db := db.New()
	defer db.Close()

	err = scripts.FetchAndRegisterMcdonaldsMenu(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("----------batch END----------")
}
