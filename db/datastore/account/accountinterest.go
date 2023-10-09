package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createAccountInterest = `-- name: CreateAccountInterest: one
INSERT INTO Account_Interest (
Account_Id, Interest, Effective_Rate, Interest_Rate, Credit, 
Debit, Accruals, Waived_Int, Last_Accrued_Date
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9
) 
ON CONFLICT(Account_Id) DO UPDATE SET
	Accruals = excluded.Accruals,
	Credit = excluded.Credit,
	Debit = excluded.Debit,
	Effective_Rate = excluded.Effective_Rate,
	Interest = excluded.Interest,
	Interest_Rate = excluded.Interest_Rate,
	Last_Accrued_Date = excluded.Last_Accrued_Date,
	Waived_Int = excluded.Waived_Int
RETURNING Account_Id, UUId, Interest, Effective_Rate, Interest_Rate, Credit, 
Debit, Accruals, Waived_Int, Last_Accrued_Date
`

type AccountInterestRequest struct {
	AccountId       int64           `json:"accountId"`
	Uuid            uuid.UUID       `json:"uuid"`
	Interest        decimal.Decimal `json:"interest"`
	EffectiveRate   decimal.Decimal `json:"effectiveRate"`
	InterestRate    decimal.Decimal `json:"interestRate"`
	Credit          decimal.Decimal `json:"credit"`
	Debit           decimal.Decimal `json:"debit"`
	Accruals        decimal.Decimal `json:"accruals"`
	WaivedInt       decimal.Decimal `json:"waivedInt"`
	LastAccruedDate sql.NullTime    `json:"lastAccruedDate"`
}

func (q *QueriesAccount) CreateAccountInterest(ctx context.Context, arg AccountInterestRequest) (model.AccountInterest, error) {
	row := q.db.QueryRowContext(ctx, createAccountInterest,
		arg.AccountId,
		arg.Interest,
		arg.EffectiveRate,
		arg.InterestRate,
		arg.Credit,
		arg.Debit,
		arg.Accruals,
		arg.WaivedInt,
		arg.LastAccruedDate,
	)
	var i model.AccountInterest
	err := row.Scan(
		&i.AccountId,
		&i.Uuid,
		&i.Interest,
		&i.EffectiveRate,
		&i.InterestRate,
		&i.Credit,
		&i.Debit,
		&i.Accruals,
		&i.WaivedInt,
		&i.LastAccruedDate,
	)
	return i, err
}

const deleteAccountInterest = `-- name: DeleteAccountInterest :exec
DELETE FROM Account_Interest
WHERE account_id = $1
`

func (q *QueriesAccount) DeleteAccountInterest(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccountInterest, id)
	return err
}

type AccountInterestInfo struct {
	AccountId       int64           `json:"accountId"`
	Uuid            uuid.UUID       `json:"uuid"`
	Interest        decimal.Decimal `json:"interest"`
	EffectiveRate   decimal.Decimal `json:"effectiveRate"`
	InterestRate    decimal.Decimal `json:"interestRate"`
	Credit          decimal.Decimal `json:"credit"`
	Debit           decimal.Decimal `json:"debit"`
	Accruals        decimal.Decimal `json:"accruals"`
	WaivedInt       decimal.Decimal `json:"waivedInt"`
	LastAccruedDate sql.NullTime    `json:"lastAccruedDate"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getAccountInterest = `-- name: GetAccountInterest :one
SELECT 
Account_Id, mr.UUId, Interest, Effective_Rate, Interest_Rate, 
Credit, Debit, Accruals, Waived_Int, Last_Accrued_Date
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Interest d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Account_Id = $1 LIMIT 1
`

func (q *QueriesAccount) GetAccountInterest(ctx context.Context, id int64) (AccountInterestInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccountInterest, id)
	var i AccountInterestInfo
	err := row.Scan(
		&i.AccountId,
		&i.Uuid,
		&i.Interest,
		&i.EffectiveRate,
		&i.InterestRate,
		&i.Credit,
		&i.Debit,
		&i.Accruals,
		&i.WaivedInt,
		&i.LastAccruedDate,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAccountInterestbyUuid = `-- name: GetAccountInterestbyUuid :one
SELECT 
Account_Id, mr.UUId, Interest, Effective_Rate, Interest_Rate, 
Credit, Debit, Accruals, Waived_Int, Last_Accrued_Date
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Interest d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccount) GetAccountInterestbyUuid(ctx context.Context, uuid uuid.UUID) (AccountInterestInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccountInterestbyUuid, uuid)
	var i AccountInterestInfo
	err := row.Scan(
		&i.AccountId,
		&i.Uuid,
		&i.Interest,
		&i.EffectiveRate,
		&i.InterestRate,
		&i.Credit,
		&i.Debit,
		&i.Accruals,
		&i.WaivedInt,
		&i.LastAccruedDate,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listAccountInterest = `-- name: ListAccountInterest:many
SELECT 
	Account_Id, mr.UUId, Interest, Effective_Rate, Interest_Rate, 
	Credit, Debit, Accruals, Waived_Int, Last_Accrued_Date,
	mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Interest d INNER JOIN Main_Record mr on mr.UUId = d.UUId
ORDER BY uuid
LIMIT $1
OFFSET $2
`

type ListAccountInterestParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountInterest(ctx context.Context, arg ListAccountInterestParams) ([]AccountInterestInfo, error) {
	rows, err := q.db.QueryContext(ctx, listAccountInterest, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountInterestInfo{}
	for rows.Next() {
		var i AccountInterestInfo
		if err := rows.Scan(
			&i.AccountId,
			&i.Uuid,
			&i.Interest,
			&i.EffectiveRate,
			&i.InterestRate,
			&i.Credit,
			&i.Debit,
			&i.Accruals,
			&i.WaivedInt,
			&i.LastAccruedDate,

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

const updateAccountInterest = `-- name: UpdateAccountInterest :one
UPDATE Account_Interest SET 
	UUId = $2,
	Interest = $3,
	Effective_Rate = $4,
	Interest_Rate = $5,
	Credit = $6,
	Debit = $7,
	Accruals = $8,
	Waived_Int = $9,
	Last_Accrued_Date = $10
WHERE Account_Id = $1
RETURNING Account_Id, UUId, Interest, Effective_Rate, Interest_Rate, Credit, 
Debit, Accruals, Waived_Int, Last_Accrued_Date
`

func (q *QueriesAccount) UpdateAccountInterest(ctx context.Context, arg AccountInterestRequest) (model.AccountInterest, error) {
	row := q.db.QueryRowContext(ctx, updateAccountInterest,

		arg.AccountId,
		arg.Uuid,
		arg.Interest,
		arg.EffectiveRate,
		arg.InterestRate,
		arg.Credit,
		arg.Debit,
		arg.Accruals,
		arg.WaivedInt,
		arg.LastAccruedDate,
	)
	var i model.AccountInterest
	err := row.Scan(
		&i.AccountId,
		&i.Uuid,
		&i.Interest,
		&i.EffectiveRate,
		&i.InterestRate,
		&i.Credit,
		&i.Debit,
		&i.Accruals,
		&i.WaivedInt,
		&i.LastAccruedDate,
	)
	return i, err
}
