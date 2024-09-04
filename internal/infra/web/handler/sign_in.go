package handler

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/internal/enum"
	"github.com/wesleyfebarretos/challenge-bravo/internal/exception"
	"github.com/wesleyfebarretos/challenge-bravo/internal/usecase"
)

type SignInHandler struct {
	useCase usecase.SignInUseCase
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r SignInRequest) MapToDomain() usecase.SignInUseCaseParamsDto {
	return usecase.SignInUseCaseParamsDto{
		Email:    r.Email,
		Password: r.Password,
	}
}

func (r SignInRequest) Valid() error {
	reqErrors := []string{}
	if r.Email == "" {
		reqErrors = append(reqErrors, "email is required")
	}

	if r.Password == "" {
		reqErrors = append(reqErrors, "password is required")
	}

	if len(reqErrors) > 0 {
		return errors.New(strings.Join(reqErrors, ", "))
	}

	return nil
}

type SignInResponse struct {
	User  SignInUserResponse `json:"user"`
	Token string             `json:"token"`
}

type SignInUserResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Active    *bool     `json:"active"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Role      enum.Role `json:"role"`
	ID        int       `json:"id"`
}

func (r SignInResponse) MapToResponse(p usecase.SignInUseCaseResponseDto) SignInResponse {
	return SignInResponse{
		User: SignInUserResponse{
			CreatedAt: p.User.CreatedAt,
			UpdatedAt: p.User.UpdatedAt,
			Active:    p.User.Active,
			FirstName: p.User.FirstName,
			LastName:  p.User.LastName,
			Email:     p.User.Email,
			Role:      p.User.Role,
			ID:        p.User.ID,
		},
		Token: p.Token,
	}
}

func (h SignInHandler) Execute(c *gin.Context) {
	body := SignInRequest{}

	readBody(c, &body)

	err := body.Valid()
	if err != nil {
		panic(exception.BadRequest(err.Error()))
	}

	signIn := h.useCase.Execute(c, body.MapToDomain())

	res := SignInResponse{}

	c.JSON(http.StatusOK, res.MapToResponse(signIn))
}

func NewSignInHandler(useCase usecase.SignInUseCase) SignInHandler {
	return SignInHandler{
		useCase: useCase,
	}
}
