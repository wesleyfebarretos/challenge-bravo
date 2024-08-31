package exception

import "net/http"

type NotFoundException struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"not found"`
}

func (e *NotFoundException) Error() string {
	return e.Message
}

func NotFound(message string) *NotFoundException {
	return &NotFoundException{
		Code:    http.StatusNotFound,
		Message: message,
	}
}
