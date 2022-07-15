package customerror

import "net/http"

var (
	ErrUnprocessableEntity = NewError(http.StatusUnprocessableEntity, "Unprocessable entity")
	ErrNotFound            = NewError(http.StatusNotFound, "Not found")
	ErrInternalServerError = NewError(http.StatusInternalServerError, "Internal Server Error")
	ErrBadRequest          = NewError(http.StatusBadRequest, "Bad Request")
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) NewMsg(msg string) *Error {
	return NewError(e.Code, msg)
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
