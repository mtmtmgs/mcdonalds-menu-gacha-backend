package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	Base
	LastName  string `bun:"last_name,notnull"`
	FirstName string `bun:"first_name,notnull"`
	Email     string `bun:"email,unique,notnull"`
	Password  string `bun:"password,notnull"`
}
