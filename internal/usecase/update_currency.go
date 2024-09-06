package usecase

import (
	"context"
	"time"

	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/exception"
)

type UpdateCurrencyUseCase struct {
	repository entity.CurrencyRepository
}

func (u UpdateCurrencyUseCase) Execute(c context.Context, p entity.Currency, id, userID int) {
	p.ID = id
	p.UpdatedAt = time.Now().UTC()
	p.UpdatedBy = userID

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
