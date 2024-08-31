package usecase

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/exception"
)

type UpdateUserUseCase struct {
	repository entity.UserRepository
}

func (u UpdateUserUseCase) Execute(c context.Context, p entity.User) {
	err := u.repository.Update(c, p)
	if err != nil {
		panic(exception.InternalServer(err.Error()))
	}
}

func NewUpdateUserUseCase(repository entity.UserRepository) UpdateUserUseCase {
	return UpdateUserUseCase{
		repository: repository,
	}
}
