package main

import (
	"log"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/batch/scripts"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/db"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/env"
)

func main() {
	log.Println("------------------------------------------------------------")
	log.Println("BEGIN batch")
	log.Println("------------------------------------------------------------")
	env := env.NewCmdEnv("../.env")
	db := db.New(env)
	defer db.Close()

	err := scripts.FetchAndRegisterMcdonaldsMenu(db)
	if err != nil {
		log.Println("------------------------------------------------------------")
		log.Println("EXECUTE FAILURE!")
		log.Println("------------------------------------------------------------")
		log.Fatalf("DETAIL: %v", err)
	}
	log.Println("------------------------------------------------------------")
	log.Println("FINISHED SUCCESSFUL batch")
	log.Println("------------------------------------------------------------")
}
