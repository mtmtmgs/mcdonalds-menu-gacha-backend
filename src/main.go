package main

import (
	"os"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/controllers"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/db"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/env"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/repositories"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/router"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/services"
)

func main() {
	env := env.NewAppEnv()

	db := db.New(env)
	repositories := repositories.New(db)
	services := services.New(repositories)
	controllers := controllers.New(services)

	r := router.New(env)
	router.Register(r, controllers)

	r.Logger.Fatal(r.Start(":" + os.Getenv("WEB_PORT")))
}
