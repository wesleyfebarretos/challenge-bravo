package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type testTxWrapper struct {
	tx     pgx.Tx
	closed bool
}

var TestTxWrapper = &testTxWrapper{}

func (t *testTxWrapper) IsClosed() bool {
	return t.closed
}

func (t *testTxWrapper) Close(ctx context.Context) {
	if t.tx == nil {
		t.closed = true
		return
	}

	if err := t.tx.Rollback(ctx); err != nil {
		log.Fatalf("error on closing transaction in test enviroment: %v", err)
	}
	t.closed = true
	t.tx = nil
}

func BeginTestTxWrapper(ctx context.Context) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Fatalf("error on opening transaction to test enviroment: %v", err)
	}

	TestTxWrapper = &testTxWrapper{
		tx:     tx,
		closed: false,
	}
}
