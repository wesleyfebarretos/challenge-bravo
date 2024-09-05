package currency_repository

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/repository/sqlc/currency_connection"
)

func FindAllMapToEntity(p []currency_connection.Currency) []entity.Currency {
	res := []entity.Currency{}

	for _, v := range p {
		res = append(res, entity.Currency{
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
			CountryCode:     v.CountryCode,
			Number:          v.Number,
			SearchURL:       v.SearchUrl,
			Fic:             &v.Fic,
			Country:         v.Country,
			Name:            v.Name,
			Code:            v.Code,
			USDExchangeRate: v.UsdExchangeRate,
			ID:              v.ID,
			CreatedBy:       v.CreatedBy,
			UpdatedBy:       *v.UpdatedBy,
			RealTimeValue:   v.RealTimeValue,
		})
	}

	return res
}

func (c CurrencyRepository) FindAll(ctx context.Context) ([]entity.Currency, error) {
	currencies, err := c.queries.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return FindAllMapToEntity(currencies), nil
}
