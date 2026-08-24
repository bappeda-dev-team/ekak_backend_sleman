package web

// CustomError untuk menandai error yang harus dikembalikan sebagai Bad Request
type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewBadRequestError(message string) *CustomError {
	return &CustomError{
		Code:    400,
		Message: message,
	}
}

func NewNotFoundError(message string) *CustomError {
	return &CustomError{
		Code:    404,
		Message: message,
	}
}
