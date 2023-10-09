package db

import (
	"database/sql"
	common "simplebank/db/common"
)

type TestOtherInfo struct {
	Greet string
	Name  string
}

func New(db common.DBTX) *QueriesIdentity {
	return &QueriesIdentity{db: db}
}

type QueriesIdentity struct {
	db common.DBTX
	// DB  *sql.DB
	// Row *sql.Row
}

func (q *QueriesIdentity) WithTx(tx *sql.Tx) *QueriesIdentity {
	return &QueriesIdentity{
		db: tx,
	}
}
