package db

import (
	"log"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/env"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func New(env env.Env) *bun.DB {
	db := bun.NewDB(env.DBConnect, pgdialect.New())
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db.AddQueryHook(env.QueryHook)
	return db
}
