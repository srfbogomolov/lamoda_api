package sql

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/srfbogomolov/warehouse_api/internal/domain/sql/sqlstore"
)

var errInvalidTxType = errors.New("invalid tx type, tx type should be *sqlx.Tx")

type ctxTransactionKey struct{}

type sqlRepository interface {
	getDB() *sqlx.DB
}

func getSqlxDatabase(ctx context.Context, r sqlRepository) (sqlstore.SqlxDatabase, error) {
	txv := ctx.Value(ctxTransactionKey{})
	if txv == nil {
		return r.getDB(), nil
	}
	if tx, ok := txv.(*sqlx.Tx); ok {
		return tx, nil
	}
	return nil, errInvalidTxType
}

func inSqlTransaction(ctx context.Context, r sqlRepository, fn func(context.Context) error) error {
	tx, err := r.getDB().BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	trxCtx := context.WithValue(ctx, ctxTransactionKey{}, tx)

	err = fn(trxCtx)
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
