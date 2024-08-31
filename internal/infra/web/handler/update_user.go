package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/exception"
	"github.com/wesleyfebarretos/challenge-bravo/internal/usecase"
)

type UpdateUserHandler struct {
	useCase usecase.UpdateUserUseCase
}

type UpdateUserRequest struct {
	Active    *bool  `json:"active,omitempty" example:"true"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Password  string `json:"password" example:"12$a@3$@00!"`
	Email     string `json:"email" example:"johndoe@gmail.com"`
	ID        int    `json:"id"`
}

func (h UpdateUserRequest) MapToDomain() entity.User {
	return entity.User{
		Active:    h.Active,
		FirstName: h.FirstName,
		LastName:  h.LastName,
		Password:  h.Password,
		Email:     h.Email,
	}
}

func (h UpdateUserRequest) Validate() {
	if h.FirstName == "" {
		panic(exception.BadRequest("first_name is required"))
	}

	if h.LastName == "" {
		panic(exception.BadRequest("last_name is required"))
	}
	if h.Password == "" {
		panic(exception.BadRequest("password is required"))
	}
	if h.Email == "" {
		panic(exception.BadRequest("email is required"))
	}
}

func (h UpdateUserHandler) Execute(c *gin.Context) {
	body := UpdateUserRequest{}

	readBody(c, &body)

	h.useCase.Execute(c, body.MapToDomain())

	c.JSON(http.StatusOK, true)
}
