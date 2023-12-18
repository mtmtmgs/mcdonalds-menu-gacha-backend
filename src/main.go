package main

import (
	"log"
	"os"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// d := db.New()

	r := router.New()
	r.Logger.Fatal(r.Start(":" + os.Getenv("WEB_PORT")))
}
