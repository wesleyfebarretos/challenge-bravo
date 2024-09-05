package currency_repository

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/repository/sqlc/currency_connection"
)

func UpdateMapToDB(p entity.Currency) currency_connection.UpdateParams {
	return currency_connection.UpdateParams{
		Name:          p.Name,
		Code:          p.Code,
		RealTimeValue: p.RealTimeValue,
		UpdatedBy:     &p.UpdatedBy,
		Number:        p.Number,
		Country:       p.Country,
		CountryCode:   p.CountryCode,
		SearchUrl:     p.SearchURL,
	}
}

func (c CurrencyRepository) Update(ctx context.Context, p entity.Currency) error {
	err := c.queries.Update(ctx, UpdateMapToDB(p))
	if err != nil {
		return err
	}

	return nil
}
