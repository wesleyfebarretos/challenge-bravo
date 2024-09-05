package usecase

import (
	"context"
	"strings"

	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/exception"
)

type CreateCurrencyUseCase struct {
	repository entity.CurrencyRepository
}

func (u CreateCurrencyUseCase) Execute(c context.Context, p entity.Currency, userID int) entity.Currency {
	p.CreatedBy = userID
	p.UpdatedBy = userID

	p.Code = strings.ToUpper(p.Code)

	if p.CountryCode != nil {
		toUpper := strings.ToUpper(*p.CountryCode)
		p.CountryCode = &toUpper
	}

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
