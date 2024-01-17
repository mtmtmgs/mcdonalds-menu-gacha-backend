package entities

import (
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/models"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/domains/values"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
	"github.com/pkg/errors"
)

/*
新規インスタンス生成
新規追加の際はこれを使う
*/
func NewUser(
	lastName values.UserLastName,
	firstName values.UserFirstName,
	email values.Email,
	password values.UserPassword,
) (models.User, error) {
	var user models.User
	err := lastName.Validate()
	if err != nil {
		return user, err
	}
	err = firstName.Validate()
	if err != nil {
		return user, err
	}
	err = email.Validate()
	if err != nil {
		return user, err
	}
	err = password.Validate()
	if err != nil {
		return user, err
	}
	hashedPassword, err := utils.HashPassword(password.Value())
	if err != nil {
		return user, errors.Errorf("Something went wrong")
	}

	user.LastName = lastName.Value()
	user.FirstName = firstName.Value()
	user.Email = email.Value()
	user.Password = hashedPassword
	return user, nil
}
