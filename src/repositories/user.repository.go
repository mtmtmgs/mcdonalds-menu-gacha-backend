package repositories

import (
	"context"

	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/models"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
	"github.com/uptrace/bun"
)

type IUserRepository interface {
	GetUserByEmail(string) (models.User, error)
	CreateUser(models.User) error
}

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	userRepository := UserRepository{db: db}
	utils.CheckDependencies(userRepository)
	return &userRepository
}

func (userRepository *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	ctx := context.Background()
	count, err := userRepository.db.NewSelect().Model(&user).Where("email = ?", email).ScanAndCount(ctx)
	if count == 0 {
		err = nil
	}
	return user, err
}

func (userRepository *UserRepository) CreateUser(user models.User) error {
	ctx := context.Background()
	_, err := userRepository.db.NewInsert().Model(&user).Exec(ctx)
	return err
}
