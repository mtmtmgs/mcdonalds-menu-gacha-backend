package values

import (
	"strings"

	"github.com/pkg/errors"
)

/*
ユーザ 姓
*/
type UserLastName struct {
	value string
}

func NewUserLastName(value string) UserLastName {
	return UserLastName{value: value}
}

func (lastName *UserLastName) Validate() error {
	return nil
}

func (lastName *UserLastName) Value() string {
	return lastName.value
}

/*
ユーザ 名
*/
type UserFirstName struct {
	value string
}

func NewUserFirstName(value string) UserFirstName {
	return UserFirstName{value: value}
}

func (ufn *UserFirstName) Validate() error {
	return nil
}

func (ufn *UserFirstName) Value() string {
	return ufn.value
}

/*
ユーザ パスワード
*/
type UserPassword struct {
	value string
}

func NewUserPassword(value string) UserPassword {
	return UserPassword{value: value}
}

func (up *UserPassword) Validate() error {
	if len(up.value) < 6 {
		return errors.Errorf("パスワードは6文字以上の入力が必要です")
	}
	if up.value == strings.ToLower(up.value) {
		return errors.Errorf("パスワードには少なくとも1つの大文字を使用してください")
	}
	return nil
}

func (up *UserPassword) Value() string {
	return up.value
}
