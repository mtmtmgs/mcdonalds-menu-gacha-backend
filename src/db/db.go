package db

import (
	"database/sql"
	"log"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/env"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func New(env env.Env) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(env.Dsn),
	))
	db := bun.NewDB(sqldb, pgdialect.New())
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db.AddQueryHook(env.QueryHook)
	return db
}
