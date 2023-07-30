package domains

import "strconv"

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
	Code   string      `json:"code"`
}

func (r *Response) HttpCode() int {
	code, _ := strconv.Atoi(r.Code)

	return code
}
