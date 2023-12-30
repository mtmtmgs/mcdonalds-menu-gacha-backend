package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

/*
カスタムバリデーション
requestsパッケージのバリデーションタグで使用する
*/
func RegisterCustomValidation(v *validator.Validate) {
	v.RegisterValidation("passwordRule", passwordRule)
	v.RegisterValidation("dateRule", dateRule)
}

func passwordRule(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	reg := regexp.MustCompile(`[A-Z]`)
	return reg.MatchString(password)
}

func dateRule(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	return err == nil
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
		case "min":
			switch err.Type().String() {
			case "string":
				errorMessage = fmt.Sprintf("%sは%s文字以上の入力が必要です", strings.ToLower(err.Field()), err.Param())
			default:
				errorMessage = fmt.Sprintf("%sは%s以上の入力が必要です", strings.ToLower(err.Field()), err.Param())
			}
		case "max":
			switch err.Type().String() {
			case "string":
				errorMessage = fmt.Sprintf("%sは%s文字以下の入力が必要です", strings.ToLower(err.Field()), err.Param())
			default:
				errorMessage = fmt.Sprintf("%sは%s以下の入力が必要です", strings.ToLower(err.Field()), err.Param())
			}
		case "excludesall":
			errorMessage = fmt.Sprintf("%sには%sの文字を使用できません", strings.ToLower(err.Field()), err.Param())
		case "containsany":
			errorMessage = fmt.Sprintf("%sには%sの文字を少なくとも1つ使用してください", strings.ToLower(err.Field()), err.Param())
		// 以下、カスタムバリデーション
		case "passwordRule":
			errorMessage = fmt.Sprintf("%sには少なくとも1つの大文字を使用してください", strings.ToLower(err.Field()))
		case "dateRule":
			errorMessage = fmt.Sprintf("%sはyyyy-MM-dd形式の入力が必要です", strings.ToLower(err.Field()))
		default:
			errorMessage = fmt.Sprintf("%s", err)
		}
		errorMessages = append(errorMessages, errorMessage)
	}
	err := errors.New(strings.Join(errorMessages, "\n"))
	return err
}
