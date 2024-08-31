package user_repository

import (
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/db"
	"github.com/wesleyfebarretos/challenge-bravo/internal/infra/repository/sqlc/user_connection"
)

var (
	once       sync.Once
	repository *UserRepository
)

type UserRepository struct {
	queries *user_connection.Queries
}

func (u UserRepository) WithTx(tx pgx.Tx) UserRepository {
	return UserRepository{
		queries: u.queries.WithTx(tx),
	}
}

func (u UserRepository) New() *UserRepository {
	once.Do(func() {
		repository = &UserRepository{
			queries: user_connection.New(db.Conn),
		}
	})
	return repository
}
