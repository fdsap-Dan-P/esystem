package db

import (
	"database/sql"
	common "simplebank/db/common"
)

func New(db common.DBTX) *QueriesLocal {
	return &QueriesLocal{db: db}
}

type QueriesLocal struct {
	db common.DBTX
}

func (q *QueriesLocal) WithTx(tx *sql.Tx) *QueriesLocal {
	return &QueriesLocal{
		db: tx,
	}
}
