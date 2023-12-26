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
		models := []interface{}{
			(*models.User)(nil),
			(*models.Menu)(nil),
		}
		for _, model := range models {
			_, err := db.NewCreateTable().
				Model(model).
				IfNotExists().
				Exec(ctx)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")
		models := []interface{}{
			(*models.User)(nil),
			(*models.Menu)(nil),
		}
		for _, model := range models {
			_, err := db.NewDropTable().
				Model(model).
				IfExists().
				Exec(ctx)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})
}
