package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createAccountTerm = `-- name: CreateAccountTerm: one
INSERT INTO Account_Term (
Account_Id, Frequency, N, Paid_N, Fixed_Due, 
Cummulative_Due, Date_Start, Maturity
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8
) 
ON CONFLICT(Account_Id) DO UPDATE SET
	frequency = excluded.frequency,
	n = excluded.n,
	Paid_N = excluded.Paid_N,
	Fixed_Due = excluded.Fixed_Due,
	Cummulative_Due = excluded.Cummulative_Due,
	Date_Start = excluded.Date_Start,
	Maturity = excluded.Maturity
RETURNING Account_Id, UUId, Frequency, N, Paid_N, Fixed_Due, 
Cummulative_Due, Date_Start, Maturity
`

type AccountTermRequest struct {
	AccountId      int64           `json:"accountId"`
	Uuid           uuid.UUID       `json:"uuid"`
	Frequency      int16           `json:"frequency"`
	N              int16           `json:"n"`
	PaidN          int16           `json:"paidN"`
	FixedDue       decimal.Decimal `json:"fixedDue"`
	CummulativeDue decimal.Decimal `json:"cummulativeDue"`
	DateStart      time.Time       `json:"dateStart"`
	Maturity       time.Time       `json:"maturity"`
}

func (q *QueriesAccount) CreateAccountTerm(ctx context.Context, arg AccountTermRequest) (model.AccountTerm, error) {
	row := q.db.QueryRowContext(ctx, createAccountTerm,
		arg.AccountId,
		arg.Frequency,
		arg.N,
		arg.PaidN,
		arg.FixedDue,
		arg.CummulativeDue,
		arg.DateStart,
		arg.Maturity,
	)
	var i model.AccountTerm
	err := row.Scan(
		&i.AccountId,
		&i.Uuid,
		&i.Frequency,
		&i.N,
		&i.PaidN,
		&i.FixedDue,
		&i.CummulativeDue,
		&i.DateStart,
		&i.Maturity,
	)
	return i, err
}

const deleteAccountTerm = `-- name: DeleteAccountTerm :exec
DELETE FROM Account_Term
WHERE account_id = $1
`

func (q *QueriesAccount) DeleteAccountTerm(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccountTerm, id)
	return err
}

type AccountTermInfo struct {
	AccountId      int64           `json:"accountId"`
	Uuid           uuid.UUID       `json:"uuid"`
	Frequency      int16           `json:"frequency"`
	N              int16           `json:"n"`
	PaidN          int16           `json:"paidN"`
	FixedDue       decimal.Decimal `json:"fixedDue"`
	CummulativeDue decimal.Decimal `json:"cummulativeDue"`
	DateStart      time.Time       `json:"dateStart"`
	Maturity       time.Time       `json:"maturity"`
	ModCtr         int64           `json:"modCtr"`
	Created        sql.NullTime    `json:"created"`
	Updated        sql.NullTime    `json:"updated"`
}

const getAccountTerm = `-- name: GetAccountTerm :one
SELECT 
Account_Id, mr.UUId, Frequency, N, 
Paid_N, Fixed_Due, Cummulative_Due, Date_Start, Maturity
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Term d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE account_id = $1 LIMIT 1
`

func (q *QueriesAccount) GetAccountTerm(ctx context.Context, id int64) (AccountTermInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccountTerm, id)
	var i AccountTermInfo
	err := row.Scan(
		&i.AccountId,
		&i.Uuid,
		&i.Frequency,
		&i.N,
		&i.PaidN,
		&i.FixedDue,
		&i.CummulativeDue,
		&i.DateStart,
		&i.Maturity,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAccountTermbyUuid = `-- name: GetAccountTermbyUuid :one
SELECT 
Account_Id, mr.UUId, Frequency, N, 
Paid_N, Fixed_Due, Cummulative_Due, Date_Start, Maturity
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Term d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccount) GetAccountTermbyUuid(ctx context.Context, uuid uuid.UUID) (AccountTermInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccountTermbyUuid, uuid)
	var i AccountTermInfo
	err := row.Scan(
		&i.AccountId,
		&i.Uuid,
		&i.Frequency,
		&i.N,
		&i.PaidN,
		&i.FixedDue,
		&i.CummulativeDue,
		&i.DateStart,
		&i.Maturity,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listAccountTerm = `-- name: ListAccountTerm:many
SELECT 
Account_Id, mr.UUId, Frequency, N, 
Paid_N, Fixed_Due, Cummulative_Due, Date_Start, Maturity
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Term d INNER JOIN Main_Record mr on mr.UUId = d.UUId
LIMIT $1
OFFSET $2
`

type ListAccountTermParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountTerm(ctx context.Context, arg ListAccountTermParams) ([]AccountTermInfo, error) {
	rows, err := q.db.QueryContext(ctx, listAccountTerm, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountTermInfo{}
	for rows.Next() {
		var i AccountTermInfo
		if err := rows.Scan(
			&i.AccountId,
			&i.Uuid,
			&i.Frequency,
			&i.N,
			&i.PaidN,
			&i.FixedDue,
			&i.CummulativeDue,
			&i.DateStart,
			&i.Maturity,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccountTerm = `-- name: UpdateAccountTerm :one
UPDATE Account_Term SET 
	Frequency = $2,
	N = $3,
	Paid_N = $4,
	Fixed_Due = $5,
	Cummulative_Due = $6,
	Date_Start = $7,
	Maturity = $8
WHERE Account_Id = $1
RETURNING Account_Id, UUId, Frequency, N, Paid_N, Fixed_Due, 
Cummulative_Due, Date_Start, Maturity
`

func (q *QueriesAccount) UpdateAccountTerm(ctx context.Context, arg AccountTermRequest) (model.AccountTerm, error) {
	row := q.db.QueryRowContext(ctx, updateAccountTerm,
		arg.AccountId,
		arg.Frequency,
		arg.N,
		arg.PaidN,
		arg.FixedDue,
		arg.CummulativeDue,
		arg.DateStart,
		arg.Maturity,
	)
	var i model.AccountTerm
	err := row.Scan(
		&i.AccountId,
		&i.Uuid,
		&i.Frequency,
		&i.N,
		&i.PaidN,
		&i.FixedDue,
		&i.CummulativeDue,
		&i.DateStart,
		&i.Maturity,
	)
	return i, err
}
