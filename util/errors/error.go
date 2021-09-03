package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ErrorTypeUnprocessaableEntity gin.ErrorType = 1 << 13
)

var (
	ErrTetapTenangTetapSemangat = CustomError{
		Message:  "Tetap Tenang Tetap Semangat",
		HTTPCode: http.StatusInternalServerError,
	}

	ErrUnauthorized = CustomError{
		Message:  "Unauthorized",
		HTTPCode: http.StatusUnauthorized,
	}

	ErrForbidden = CustomError{
		Message:  "Forbidden",
		HTTPCode: http.StatusForbidden,
	}

	ErrNotFound = CustomError{
		Message:  "Record not exist",
		HTTPCode: http.StatusNotFound,
	}

	ErrUnprocessableEntity = CustomError{
		Message:  "Unprocessable Entity",
		HTTPCode: http.StatusUnprocessableEntity,
	}
)

type CustomError struct {
	Message  string `json:"message,"`
	HTTPCode int    `json:"code"`
}

func (c CustomError) Error() string {
	return c.Message
}
