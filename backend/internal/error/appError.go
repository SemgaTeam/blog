package error 

import (
	"fmt"
	"errors"
	"net/http"
)

type AppError struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Err error `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func BadRequest(err error, msg string) *AppError {
	return &AppError{
		Code: http.StatusBadRequest, 
		Message: msg, 
		Err: err,
	}
}

func Unauthorized(err error, msg string) *AppError {
	return &AppError{
		Code: http.StatusUnauthorized, 
		Message: msg, 
		Err: err,
	}
}

func Forbidden(err error, msg string) *AppError {
	return &AppError{
		Code: http.StatusForbidden, 
		Message: msg, 
		Err: err,
	}
}

func NotFound(err error, msg string) *AppError {
	return &AppError{
		Code: http.StatusNotFound, 
		Message: msg, 
		Err: err,
	}
}

func Internal(err error) *AppError {
	return &AppError{
		Code: http.StatusInternalServerError, 
		Message: "internal server error", 
		Err: err,
	}
}

func FromError(err error) *AppError {
	var appErr *AppError

	if errors.As(err, &appErr) {
		return appErr
	}

	return Internal(err)
}
