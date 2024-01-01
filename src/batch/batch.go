package main

import (
	"log"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/batch/scripts"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/db"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("------------------------------------------------------------")
	log.Println("BEGIN batch")
	log.Println("------------------------------------------------------------")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("------------------------------------------------------------")
		log.Println("EXECUTE FAILURE!")
		log.Println("------------------------------------------------------------")
		log.Fatalf("DETAIL: %v", err)
	}

	db := db.New()
	defer db.Close()

	err = scripts.FetchAndRegisterMcdonaldsMenu(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("------------------------------------------------------------")
	log.Println("FINISHED SUCCESSFUL batch")
	log.Println("------------------------------------------------------------")
}
