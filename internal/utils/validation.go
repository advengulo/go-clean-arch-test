package utils

func validationTagMessage(tag string) string {
	switch tag {
	case "required":
		return "required field"
	case "email":
		return "invalid email"
	case "min":
		return "error minimum"
	}
	return ""
}
