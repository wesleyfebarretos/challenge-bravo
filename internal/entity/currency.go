package entity

import "context"

type Currency struct {
	Country         *string
	Number          *int
	SearchURL       *string
	Fic             *bool
	Code            string
	Name            string
	RealTimeValue   float64
	USDExchangeRate float64
	ID              int
}

type CurrencyRepository interface {
	Create(context.Context, Currency) (Currency, error)
	FindOneById(context.Context, int32) (*Currency, error)
	FindOneByCode(context.Context, string) (*Currency, error)
	Update(context.Context, Currency) error
	Delete(context.Context, int32) error
	FindAll(context.Context) ([]Currency, error)
}
