package validators

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

func notRussian(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	for _, r := range str {
		if unicode.Is(unicode.Cyrillic, r) {
			return false
		}
	}
	return true
}

func RegisterValidators(v *validator.Validate) {
	v.RegisterValidation("not_russian", notRussian)
}
