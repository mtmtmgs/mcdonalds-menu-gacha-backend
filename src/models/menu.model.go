package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Menu struct {
	bun.BaseModel `bun:"table:menus,alias:m"`

	ID        int64     `bun:"id,pk,autoincrement"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	Name      string    `bun:"name,notnull"`
	Price     int64     `bun:"price,nullzero"`
	Category  string    `bun:"category"`
	Hourly    string    `bun:"hourly"`
}
