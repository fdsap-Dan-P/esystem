package db

import (
	"database/sql"
	dsCommon "simplebank/db/common"
)

func New(db dsCommon.DBTX) *QueriesTransaction {
	return &QueriesTransaction{db: db}
}

type QueriesTransaction struct {
	db dsCommon.DBTX
}

func (q *QueriesTransaction) WithTx(tx *sql.Tx) *QueriesTransaction {
	return &QueriesTransaction{
		db: tx,
	}
}
