package currency_repository

import (
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/config"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/entity"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/enum"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/infra/db"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/infra/repository/sqlc/currency_connection"
)

var (
	once       sync.Once
	repository *CurrencyRepository
)

type CurrencyRepository struct {
	queries *currency_connection.Queries
}

func (u *CurrencyRepository) RenewTestTx(tx db.DBConn) {
	u.queries = currency_connection.New(db.GetConnection())
}

func (u CurrencyRepository) WithTx(tx pgx.Tx) entity.CurrencyRepository {
	return &CurrencyRepository{
		queries: u.queries.WithTx(tx),
	}
}

func New() entity.CurrencyRepository {
	if config.Envs.AppEnv == enum.TEST_ENVIROMENT && repository != nil {
		repository.RenewTestTx(db.GetConnection())
	}

	once.Do(func() {
		repository = &CurrencyRepository{
			queries: currency_connection.New(db.GetConnection()),
		}
	})
	return repository
}
