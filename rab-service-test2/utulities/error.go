package errors

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