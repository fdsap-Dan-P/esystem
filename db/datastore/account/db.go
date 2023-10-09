package db

import (
	dsCommon "simplebank/db/common"
)

func New(db dsCommon.DBTX) *QueriesAccount {
	return &QueriesAccount{db: db}
}

type QueriesAccount struct {
	db dsCommon.DBTX
	// DB  *sql.DB
	// Row *sql.Row
}
