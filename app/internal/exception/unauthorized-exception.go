package exception

import "net/http"

type UnauthorizedException struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"bad request"`
}

func (e *UnauthorizedException) Error() string {
	return e.Message
}

func Unauthorized(message string) *HttpException {
	return &HttpException{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}
