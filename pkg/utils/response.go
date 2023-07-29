package utils

import (
	"github.com/advengulo/go-clean-arch-test/domains"
	"strconv"
)

func Response(message string, data interface{}, error interface{}, code int) (response domains.Response) {
	response.Message = message
	response.Data = data
	response.Error = error
	response.Code = strconv.Itoa(code)

	return response
}
