package utils

import (
	"fmt"
	"github.com/advengulo/go-clean-arch-test/domains"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func ErrorValidation(err validator.ValidationErrors, data interface{}) (errVal []domains.ErrorValidation) {
	t := reflect.TypeOf(data)

	for _, e := range err {
		field, found := t.FieldByName(e.Field())
		if !found {
			continue
		}

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = e.Field()
		}
		
		errVal = append(errVal, domains.ErrorValidation{
			Parameter: jsonTag,
			Message:   fmt.Sprintf("Parameter %s %s", jsonTag, validationTagMessage(e.Tag())),
		})
	}

	return
}
