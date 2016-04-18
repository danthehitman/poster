package api

import "net/http"

type apiError struct {
	Error error
	Code int
	Message string
}

func New(code int, message string, err error) *apiError {
	return &apiError{Code:code, Message:message, Error:err}
}

func NotFoundError(err error) *apiError {
	return &apiError{Code:http.StatusNotFound, Message:"Resource not found", Error: err}
}

func BadRequestError(err error) *apiError {
	return &apiError{Code:http.StatusBadRequest, Message:err.Error(), Error: err}
}

func UnauthorizedError(err error) *apiError {
	return &apiError{Code:http.StatusUnauthorized, Message:"Authorization failed", Error: err}
}

func InternalServerError(err error ) * apiError {
	return &apiError{Code:http.StatusInternalServerError, Message:"The system encountered an error", Error: err}
}