package db

import (
	"context"
	"log"

	"github.com/wesleyfebarretos/challenge-bravo/app/internal/config"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/enum"
)

func GetConnection() DBConn {
	if config.Envs.AppEnv == enum.TEST_ENVIROMENT {
		tx, err := conn.Begin(context.Background())
		if err != nil {
			log.Fatalf("error on opening transaction to test enviroment: %v", err)
		}

		conn = tx
	}

	return conn
}
