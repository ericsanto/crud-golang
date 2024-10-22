package rest_err

import (
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
}

func NewRestErr(message, err string, code int) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
	}
}

func NewBadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFound(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}
