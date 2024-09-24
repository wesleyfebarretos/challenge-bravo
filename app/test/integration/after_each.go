package integration

import (
	"context"

	"github.com/wesleyfebarretos/challenge-bravo/app/internal/infra/db"
)

func afterEach() {
	db.TestTxWrapper.Close(context.TODO())
}
