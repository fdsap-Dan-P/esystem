package db

import (
	"database/sql"
	common "simplebank/db/common"
)

func New(db common.DBTX) *QueriesDocument {
	return &QueriesDocument{db: db}
}

type QueriesDocument struct {
	db common.DBTX
}

func (q *QueriesDocument) WithTx(tx *sql.Tx) *QueriesDocument {
	return &QueriesDocument{
		db: tx,
	}
}
