package usecase

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/exception"
)

type CreateCurrencyUseCase struct {
	repository entity.CurrencyRepository
}

func (u CreateCurrencyUseCase) Execute(c context.Context, p entity.Currency) entity.Currency {
	currency, err := u.repository.Create(c, p)
	if err != nil {
		panic(exception.InternalServer(err.Error()))
	}

	return currency
}

func NewCreateCurrencyUseCase(repository entity.CurrencyRepository) CreateCurrencyUseCase {
	return CreateCurrencyUseCase{
		repository: repository,
	}
}
