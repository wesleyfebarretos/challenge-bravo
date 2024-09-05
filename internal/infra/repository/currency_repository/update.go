package currency_repository

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/repository/sqlc/currency_connection"
)

func UpdateMapToDB(p entity.Currency) currency_connection.UpdateParams {
	return currency_connection.UpdateParams{
		ID:              p.ID,
		Name:            p.Name,
		Code:            p.Code,
		Number:          p.Number,
		Country:         p.Country,
		CountryCode:     p.CountryCode,
		SearchUrl:       p.SearchURL,
		UsdExchangeRate: p.USDExchangeRate,
		Fic:             *p.Fic,
		UpdatedBy:       &p.UpdatedBy,
		UpdatedAt:       p.UpdatedAt,
	}
}

func (c CurrencyRepository) Update(ctx context.Context, p entity.Currency) error {
	err := c.queries.Update(ctx, UpdateMapToDB(p))
	if err != nil {
		return err
	}

	return nil
}
