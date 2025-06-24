package apperrors

type AppError struct {
	Message    string
	HttpStatus int
}

func (e AppError) Error() string {
	return e.Message
}
