package migrations

import (
	"context"
	"fmt"
	"log"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		_, err := db.NewCreateTable().
			Model((*models.Menu)(nil)).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		_, err := db.NewDropTable().
			Model((*models.Menu)(nil)).
			IfExists().
			Exec(ctx)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
}
