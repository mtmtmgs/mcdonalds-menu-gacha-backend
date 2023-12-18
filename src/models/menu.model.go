package models

import (
	"github.com/uptrace/bun"
)

type Menu struct {
	bun.BaseModel `bun:"table:menus,alias:m"`

	base
	Name     string `bun:"name,notnull"`
	Price    int64  `bun:"price,nullzero"`
	Category string `bun:"category"`
	Hourly   string `bun:"hourly"`
}
