package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/hm-mtmtmgs/mcdonalds-menu-gacha-backend/utils"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidatorMiddleware() *CustomValidator {
	v := validator.New()
	utils.RegisterCustomValidation(v)
	return &CustomValidator{v}
}

func (customValidator *CustomValidator) Validate(i interface{}) error {
	return customValidator.validator.Struct(i)
}
