package exception

import "net/http"

type UnauthorizedException struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"bad request"`
}

func (e *UnauthorizedException) Error() string {
	return e.Message
}

func Unauthorized(message string) *UnauthorizedException {
	return &UnauthorizedException{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}
