package errors

import "net/http"

func BadRequest(code, message string) *AppError {
	return New(code, message, http.StatusBadRequest)
}

func Unauthorized(code, message string) *AppError {
	return New(code, message, http.StatusUnauthorized)
}

func Forbidden(code, message string) *AppError {
	return New(code, message, http.StatusForbidden)
}

func NotFound(code, message string) *AppError {
	return New(code, message, http.StatusNotFound)
}

func Conflict(code, message string) *AppError {
	return New(code, message, http.StatusConflict)
}

func Internal(message string) *AppError {
	return New(CodeInternalError, message, http.StatusInternalServerError)
}
