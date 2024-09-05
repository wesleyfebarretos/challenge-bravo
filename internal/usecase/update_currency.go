package usecase

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/exception"
)

type UpdateCurrencyUseCase struct {
	repository entity.CurrencyRepository
}

func (u UpdateCurrencyUseCase) Execute(c context.Context, p entity.Currency) {
	err := u.repository.Update(c, p)
	if err != nil {
		panic(exception.InternalServer(err.Error()))
	}
}

func NewUpdateCurrencyUseCase(repository entity.CurrencyRepository) UpdateCurrencyUseCase {
	return UpdateCurrencyUseCase{
		repository: repository,
	}
}
