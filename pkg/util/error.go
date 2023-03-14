package util

type ErrorCustom struct {
	Code    int
	Message string
}

func SetError(code int, message string) error {
	return ErrorCustom{
		Code:    code,
		Message: message,
	}
}

func (ce ErrorCustom) Error() string {
	return ce.Message
}

func ErrorResponse(err error) (int, map[string]interface{}) {
	errorCustom, ok := err.(ErrorCustom)
	if !ok {
		err = SetError(500, "internal server error")
		errorCustom = err.(ErrorCustom)
	}
	return errorCustom.Code, map[string]interface{}{
		"message": errorCustom.Message,
		"status":  errorCustom.Code,
	}
}
