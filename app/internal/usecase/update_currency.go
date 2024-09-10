package usecase

import (
	"context"
	"strings"
	"time"

	"github.com/wesleyfebarretos/challenge-bravo/app/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/exception"
)

type UpdateCurrencyUseCase struct {
	repository entity.CurrencyRepository
}

func (u UpdateCurrencyUseCase) Execute(c context.Context, p entity.Currency, id, userID int) {
	p.ID = id
	p.UpdatedAt = time.Now().UTC()
	p.UpdatedBy = userID

	p.Code = strings.ToUpper(p.Code)

	if p.CountryCode != nil {
		upperCountryCode := strings.ToUpper(*p.CountryCode)
		p.CountryCode = &upperCountryCode
	}

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
