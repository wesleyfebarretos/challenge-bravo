package usecase

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/exception"
)

type FindCurrencyByCodeUseCase struct {
	repository entity.CurrencyRepository
}

func (u FindCurrencyByCodeUseCase) Execute(c context.Context, code string) *entity.Currency {
	currency, err := u.repository.FindOneByCode(c, code)
	if err != nil {
		panic(exception.InternalServer(err.Error()))
	}

	return currency
}

func NewFindCurrencyByCodeUseCase(repository entity.CurrencyRepository) FindCurrencyByCodeUseCase {
	return FindCurrencyByCodeUseCase{
		repository: repository,
	}
}
