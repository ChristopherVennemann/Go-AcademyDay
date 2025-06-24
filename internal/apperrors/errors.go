package apperrors

import (
	"net/http"
)

var (
	UserAlreadyExists = &AppError{
		Message:    "user with that name or email already exists",
		HttpStatus: http.StatusConflict,
	}
)
