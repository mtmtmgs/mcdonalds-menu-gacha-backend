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
		log.Fatal("環境変数を設定してください")
	}

	e := router.New()
	e.Logger.Fatal(e.Start(":" + os.Getenv("WEB_PORT")))
}
