package utils

import (
	"fmt"
	"github.com/advengulo/go-clean-arch-test/domains"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
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

// ErrorHandler is custom implementation of the error handler.
func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	switch e := err.(type) {
	case *echo.HTTPError:
		code = e.Code
		message = e.Message.(string)
	}

	// Send the error response.
	c.JSON(code, Response(nil, message, code))
}
