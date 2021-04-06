package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/sankester/simple_bank/utils"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return utils.IsSupportedCurrecny(currency)
	}

	return false
}
