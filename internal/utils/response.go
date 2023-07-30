package utils

import (
	"github.com/advengulo/go-clean-arch-test/domains"
	"net/http"
	"strconv"
)

func Response(data interface{}, error interface{}, code int) (response domains.Response) {
	status := "OK"
	if code != http.StatusOK {
		status = "ERROR"
	}

	response.Status = status
	response.Data = data
	response.Error = error
	response.Code = strconv.Itoa(code)

	return response
}
