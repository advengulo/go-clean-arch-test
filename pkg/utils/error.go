package utils

import (
	"fmt"
	"github.com/advengulo/go-clean-arch-test/domains"
	"github.com/go-playground/validator/v10"
)

func ErrorValidation(err validator.ValidationErrors) (errVal []domains.ErrorValidation) {
	for _, e := range err {
		errVal = append(errVal, domains.ErrorValidation{
			Field:   e.Field(),
			Message: fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", e.Field(), e.Tag()),
		})
	}

	return
}
