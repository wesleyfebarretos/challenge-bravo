package currency_repository

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/repository/sqlc/currency_connection"
)

func CreateMapToDB(p entity.Currency) currency_connection.CreateParams {
	return currency_connection.CreateParams{
		Name:          p.Name,
		Code:          p.Code,
		RealTimeValue: p.RealTimeValue,
		CreatedBy:     p.CreatedBy,
		Number:        p.Number,
		Country:       p.Country,
		CountryCode:   p.CountryCode,
		SearchUrl:     p.SearchURL,
	}
}

func CreateMapToEntity(p currency_connection.Currency) entity.Currency {
	return entity.Currency{
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
		CountryCode:   p.CountryCode,
		Number:        p.Number,
		SearchURL:     p.SearchUrl,
		Fic:           &p.Fic,
		Country:       p.Country,
		Name:          p.Name,
		Code:          p.Code,
		ID:            p.ID,
		CreatedBy:     p.CreatedBy,
		UpdatedBy:     *p.UpdatedBy,
		RealTimeValue: p.RealTimeValue,
	}
}

func (c CurrencyRepository) Create(ctx context.Context, p entity.Currency) (entity.Currency, error) {
	currency, err := c.queries.Create(ctx, CreateMapToDB(p))
	if err != nil {
		return entity.Currency{}, err
	}

	return CreateMapToEntity(currency), nil
}
