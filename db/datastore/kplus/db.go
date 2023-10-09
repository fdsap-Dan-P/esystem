package db

import (
	"context"
	"database/sql"
	common "simplebank/db/common"
	dsAcc "simplebank/db/datastore/account"
	dsRef "simplebank/db/datastore/reference"
	dsTrn "simplebank/db/datastore/transaction"
)

func New(db common.DBTX) *QueriesKPlus {
	return &QueriesKPlus{
		db:                 db,
		QueriesAccount:     dsAcc.New(db),
		QueriesReference:   dsRef.New(db),
		QueriesTransaction: dsTrn.New(db),
	}
}

type QueriesKPlus struct {
	db common.DBTX
	*dsAcc.QueriesAccount
	*dsRef.QueriesReference
	*dsTrn.QueriesTransaction
}

func (q *QueriesKPlus) WithTx(tx *sql.Tx) *QueriesKPlus {
	return &QueriesKPlus{
		db: tx,
	}
}

type QuerierKPlus interface {
	dsAcc.QuerierAccount
	dsRef.QuerierReference
	dsTrn.QuerierTransaction
	GetCustomersInfo(ctx context.Context, arr CustomersInfoParam) ([]CustomerInfo, error)
	SearchCustomerCID(ctx context.Context, cid int64) (KPLUSCustomer, error)
	SavingsList(ctx context.Context, arg SavingsListParams) ([]Savings, error)
	ColSht(ctx context.Context, arg ColShtParams) ([]ColShtperAcc, error)
	CallBackRef(ctx context.Context, prNo string) (KPLUSResponse, error)
	MultiplePayment(ctx context.Context, req MultiplePaymentRequest) (KPLUSResponse, error)
}

var _ QuerierKPlus = (*QueriesKPlus)(nil)
