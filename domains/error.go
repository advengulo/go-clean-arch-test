package domains

type ErrorValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
