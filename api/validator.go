package api

import (
	"github.com/annguyen34/simple-bank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		if util.IsSupportedCurrency(currency) {
			return true
		}
	}

	return false
}
