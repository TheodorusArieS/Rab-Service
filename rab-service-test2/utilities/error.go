package errors

import (
	"net/http"
)

type RestError struct {
	Message string `json:"message"`
	Status int64 `json:"status"`
	Error string `json:"error"`
}

func SendError(message string,status int64) *RestError{
	return &RestError{
		Message :message,
		Status : status,
		Error : "Terkena Error",
	}
}

func BadRequestError(message string) *RestError{
	return &RestError{
		Message :message,
		Status :http.StatusBadRequest,
		Error :"Bad Request",
	}
}

func NotFoundError(message string) *RestError{
	return &RestError{
		Message :message,
		Status:http.StatusNotFound,
		Error :"Not Found",
	}
}

func InternalServerError(message string) *RestError{
	return &RestError{
		Message :message,
		Status:http.StatusInternalServerError,
		Error :"Internal Service Error",
	}
}