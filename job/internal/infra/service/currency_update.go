package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/wesleyfebarretos/challenge-bravo/job/internal/infra/db"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/types"
)

type CurrencyUpdaterService struct {
	client http.Client
}

func NewCurrencyUpdaterService(client http.Client) CurrencyUpdaterService {
	return CurrencyUpdaterService{}
}

func (c CurrencyUpdaterService) GetCurrenciesExchangeRatesInUSD(ctx context.Context) ([]*types.Currency, error) {

	usd := types.Currency{}

	query := "SELECT code, search_url, response_path_to_rate FROM currency WHERE code = $1"

	db.Conn.QueryRow(ctx, query, "USD").Scan(&usd.Code, &usd.SearchURL, &usd.ResponsePathToRate)

	req, err := http.NewRequest(http.MethodGet, *usd.SearchURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	dataB, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	usdExchangeRatesMap := map[string]any{}

	err = json.Unmarshal(dataB, &usdExchangeRatesMap)
	if err != nil {
		return nil, err
	}

	usdExchangeRatesMap = usdExchangeRatesMap["rates"].(map[string]any)

	query = "SELECT id, code FROM currency WHERE search_url IS NULL"

	currencies := []*types.Currency{}

	rows, err := db.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var currency = &types.Currency{}

		err := rows.Scan(&currency.ID, &currency.Code)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	updatedCurrencies := []*types.Currency{}

	for _, v := range currencies {
		rate, ok := usdExchangeRatesMap[v.Code]
		if ok {
			v.USDExchangeRate = rate.(float64)

			updatedCurrencies = append(updatedCurrencies, v)

		}
	}

	return updatedCurrencies, nil
}

func (c CurrencyUpdaterService) UpdateRates(ctx context.Context, currencies []*types.Currency) error {
	tx, err := db.Conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, v := range currencies {
		query := `
            UPDATE currency
            SET
                usd_exchange_rate = $2,
                updated_at = $3,
                updated_by = 1
            WHERE
                id = $1`

		_, err := tx.Exec(ctx, query, v.ID, v.USDExchangeRate, time.Now().UTC())
		if err != nil {
			return err
		}
	}

	tx.Commit(ctx)

	return nil
}
