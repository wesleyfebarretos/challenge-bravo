package handler

import (
	"errors"
	"net/http"
	"strings"

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
		ID:        h.ID,
	}
}

func (h UpdateUserRequest) Valid() error {
	reqErrors := []string{}

	if h.FirstName == "" {
		reqErrors = append(reqErrors, "first_name is required")
	}

	if h.LastName == "" {
		reqErrors = append(reqErrors, "last_name is required")
	}
	if h.Password == "" {
		reqErrors = append(reqErrors, "password is required")
	}
	if h.Email == "" {
		reqErrors = append(reqErrors, "email is required")
	}

	if len(reqErrors) > 0 {
		return errors.New(strings.Join(reqErrors, ", "))
	}

	return nil
}

func (h UpdateUserHandler) Execute(c *gin.Context) {
	body := UpdateUserRequest{}

	id := getIdFromReq(c)

	readBody(c, &body)

	err := body.Valid()
	if err != nil {
		panic(exception.BadRequest(err.Error()))
	}

	body.ID = id

	h.useCase.Execute(c, body.MapToDomain())

	c.JSON(http.StatusOK, true)
}

func NewUpdateUserHandler(useCase usecase.UpdateUserUseCase) UpdateUserHandler {
	return UpdateUserHandler{
		useCase: useCase,
	}
}
