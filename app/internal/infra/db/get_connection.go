package db

import (
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/config"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/enum"
)

func GetConnection() DBConn {
	if config.Envs.AppEnv == enum.TEST_ENVIROMENT {
		return TestTxWrapper.tx
	}

	return conn
}
