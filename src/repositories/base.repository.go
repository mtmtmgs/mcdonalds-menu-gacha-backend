package repositories

import (
	"context"

	"github.com/uptrace/bun"
)

type IBaseRepository interface {
	GetDB() *bun.DB
}

type BaseRepository struct {
	db *bun.DB
}

func NewBaseRepository(db *bun.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (baseRepository *BaseRepository) GetDB() *bun.DB {
	return baseRepository.db
}

func GetById[T any](db *bun.DB, id uint) (T, error) {
	var model T
	ctx := context.Background()
	err := db.NewSelect().Model(&model).Where("id = ?", id).Scan(ctx)
	return model, err
}

func GetList[T any](db *bun.DB) ([]T, error) {
	var model []T
	ctx := context.Background()
	err := db.NewSelect().Model(&model).Scan(ctx)
	return model, err
}
