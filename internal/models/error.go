package models

type ErrorValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
