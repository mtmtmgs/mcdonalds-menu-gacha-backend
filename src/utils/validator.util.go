package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

/*
カスタムバリデーション
requestsパッケージのバリデーションタグで使用する
*/
func RegisterCustomValidation(v *validator.Validate) {
	v.RegisterValidation("passwordRule", passwordRule)
}

func passwordRule(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	re := regexp.MustCompile(`[A-Z]`)
	return re.MatchString(password)
}

/*
バリデーションエラーメッセージ
カスタムバリデーションを含みます
*/
func ConvertValidationErrorMessage(errs validator.ValidationErrors) error {
	var errorMessage string
	var errorMessages []string

	for _, err := range errs {
		switch err.Tag() {
		// 以下、標準バリデーション
		case "required":
			errorMessage = fmt.Sprintf("%sは必須入力です", strings.ToLower(err.Field()))
		// 以下、カスタムバリデーション
		case "passwordRule":
			errorMessage = fmt.Sprintf("%sには少なくとも1つの大文字を含めてください", strings.ToLower(err.Field()))
		default:
			errorMessage = fmt.Sprintf("%s", err)
		}
		errorMessages = append(errorMessages, errorMessage)
	}
	err := errors.New(strings.Join(errorMessages, "\n"))
	return err
}
