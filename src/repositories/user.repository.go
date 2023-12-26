package repositories

import (
	"context"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/models"
	"github.com/uptrace/bun"
)

type IUserRepository interface {
	GetUser() (models.User, error)
	CreateUser(models.User) error
}

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db}
}

func (userRepository *UserRepository) GetUser() (models.User, error) {
	var user models.User
	ctx := context.Background()
	err := userRepository.db.NewSelect().Model(&user).Scan(ctx)
	return user, err
}

func (userRepository *UserRepository) CreateUser(user models.User) error {
	ctx := context.Background()
	_, err := userRepository.db.NewInsert().Model(&user).Exec(ctx)
	return err
}
