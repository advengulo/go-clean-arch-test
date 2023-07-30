package domains

import "strconv"

type Response struct {
	Code   string      `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

func (r *Response) HttpCode() int {
	code, _ := strconv.Atoi(r.Code)

	return code
}
