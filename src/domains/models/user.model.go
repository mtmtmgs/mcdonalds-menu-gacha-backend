package models

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	Base
	LastName  string `bun:"last_name,notnull"`
	FirstName string `bun:"first_name,notnull"`
	Email     string `bun:"email,unique,notnull"`
	Password  string `bun:"password,notnull"`
}

func NewUser(user User) (User, error) {
	var newUser User
	lastName, err := NewUserLastName(user.LastName)
	if err != nil {
		return newUser, err
	}
	firstName, err := NewUserFirstName(user.FirstName)
	if err != nil {
		return newUser, err
	}
	email, err := NewUserEmail(user.Email)
	if err != nil {
		return newUser, err
	}

	newUser.LastName = lastName
	newUser.FirstName = firstName
	newUser.Email = email
	newUser.Password = user.Password
	return newUser, nil
}

func NewUserLastName(lastName string) (string, error) {
	return lastName, nil
}

func NewUserFirstName(firstName string) (string, error) {
	return firstName, nil
}

func NewUserEmail(email string) (string, error) {
	return email, nil
}

func NewUserPassword(password string) (string, error) {
	if len(password) < 6 {
		return "", errors.Errorf("パスワードは6文字以上の入力が必要です")
	}
	if password == strings.ToLower(password) {
		return "", errors.Errorf("パスワードには少なくとも1つの大文字を使用してください")
	}
	return password, nil
}
