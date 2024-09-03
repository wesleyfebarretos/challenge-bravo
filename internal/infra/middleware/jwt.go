package middleware

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/internal/exception"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/service"
)

func Jwt(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		panic(exception.Unauthorized("access not authorized"))
	}

	tokenString = strings.Split(tokenString, " ")[1]

	jwtService := service.NewJwtService()

	claims, err := jwtService.VerifyToken(tokenString)
	if err != nil {
		panic(exception.Unauthorized("access not authorized"))
	}

	claimsToJson, err := json.Marshal(claims)
	if err != nil {
		panic(exception.InternalServer(err.Error()))
	}

	//  TODO: Add user use case to get user by id returned

	c.Header("user", string(claimsToJson))

	c.Next()
}
