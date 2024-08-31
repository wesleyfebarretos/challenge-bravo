package usecase

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/exception"
)

type CreateUserUseCase struct {
	repository entity.UserRepository
}

func (u CreateUserUseCase) Execute(c context.Context, p entity.User) entity.User {
	user, err := u.repository.Create(c, p)
	if err != nil {
		panic(exception.InternalServer(err.Error()))
	}

	return user
}

func NewCreateUserUseCase(repository entity.UserRepository) CreateUserUseCase {
	return CreateUserUseCase{
		repository: repository,
	}
}
