package db

import (
	"database/sql"
	common "simplebank/db/common"
)

func New(db common.DBTX) *QueriesReference {
	return &QueriesReference{db: db}
}

type QueriesReference struct {
	db common.DBTX
}

func (q *QueriesReference) WithTx(tx *sql.Tx) *QueriesReference {
	return &QueriesReference{
		db: tx,
	}
}
