package integration

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/app/internal/infra/db"
)

func beforeEach() {
	if !db.TestTxWrapper.IsClosed() {
		db.TestTxWrapper.Close(context.TODO())
	}

	db.BeginTestTxWrapper(context.TODO())
}
