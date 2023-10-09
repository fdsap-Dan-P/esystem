package db

import (
	"database/sql"
	common "simplebank/db/common"
)

func New(db common.DBTX) *QueriesAccess {
	return &QueriesAccess{db: db}
}

type QueriesAccess struct {
	db common.DBTX
}

func (q *QueriesAccess) WithTx(tx *sql.Tx) *QueriesAccess {
	return &QueriesAccess{
		db: tx,
	}
}
