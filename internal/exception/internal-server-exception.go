package exception

import "net/http"

type InternalServerException struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"bad request"`
}

func (e *InternalServerException) Error() string {
	return e.Message
}

func InternalServer(message string) *InternalServerException {
	return &InternalServerException{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
