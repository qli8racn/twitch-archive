package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("datetime", dateTimeValidation("2006-01-02"))
	return v
}

// 日付の形式を検証するカスタムバリデーション
func dateTimeValidation(layout string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		_, err := time.Parse(layout, fl.Field().String())
		return err == nil
	}
}
