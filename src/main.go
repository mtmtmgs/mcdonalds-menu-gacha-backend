package main

import (
	"log"
	"os"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/db"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/router"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := db.New()
	repositories := repositories.New(db)
	services := services.New(repositories)
	controllers := controllers.New(services)

	r := router.New()
	router.Register(r, controllers)

	r.Logger.Fatal(r.Start(":" + os.Getenv("WEB_PORT")))
}
