package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createGlAccount = `-- name: CreateGlAccount: one
INSERT INTO GL_Account (
	Office_Id, COA_Id, Balance, Pending_Trn_Amt, Account_Type_Id, 
	Currency, Partition_Id, Remark, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9
) 
ON CONFLICT(Office_Id, COA_Id) DO UPDATE SET
	Office_Id  = excluded.Office_Id ,
	coa_id = excluded.coa_id,
	Balance = excluded.Balance,
	Pending_Trn_Amt = excluded.Pending_Trn_Amt,
	Currency = excluded.Currency,
	Account_Type_Id = excluded.Account_Type_Id,
	Partition_Id  = excluded.Partition_Id ,
	Remark = excluded.Remark

RETURNING Id, UUId, Office_Id, COA_Id, Balance, Pending_Trn_Amt, Account_Type_Id, 
Currency, Partition_Id, Remark, Other_Info
`

type GlAccountRequest struct {
	Id            int64           `json:"id"`
	Uuid          uuid.UUID       `json:"uuid"`
	OfficeId      int64           `json:"officeId"`
	CoaId         sql.NullInt64   `json:"coaId"`
	Balance       decimal.Decimal `json:"balance"`
	PendingTrnAmt decimal.Decimal `json:"pendingTrnAmt"`
	AccountTypeId sql.NullInt64   `json:"AccounttypeId"`
	Currency      sql.NullString  `json:"currency"`
	PartitionId   sql.NullInt64   `json:"partitionId"`
	Remark        sql.NullString  `json:"remark"`
	OtherInfo     sql.NullString  `json:"otherInfo"`
}

func (q *QueriesAccount) CreateGlAccount(ctx context.Context, arg GlAccountRequest) (model.GlAccount, error) {
	row := q.db.QueryRowContext(ctx, createGlAccount,
		arg.OfficeId,
		arg.CoaId,
		arg.Balance,
		arg.PendingTrnAmt,
		arg.AccountTypeId,
		arg.Currency,
		arg.PartitionId,
		arg.Remark,
		arg.OtherInfo,
	)
	var i model.GlAccount
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeId,
		&i.CoaId,
		&i.Balance,
		&i.PendingTrnAmt,
		&i.AccountTypeId,
		&i.Currency,
		&i.PartitionId,
		&i.Remark,
		&i.OtherInfo,
	)
	return i, err
}

const deleteGlAccount = `-- name: DeleteGlAccount :exec
DELETE FROM GL_Account
WHERE id = $1
`

func (q *QueriesAccount) DeleteGlAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteGlAccount, id)
	return err
}

type GlAccountInfo struct {
	Id            int64           `json:"id"`
	Uuid          uuid.UUID       `json:"uuid"`
	OfficeId      int64           `json:"officeId"`
	CoaId         sql.NullInt64   `json:"coaId"`
	Balance       decimal.Decimal `json:"balance"`
	PendingTrnAmt decimal.Decimal `json:"pendingTrnAmt"`
	AccountTypeId sql.NullInt64   `json:"AccounttypeId"`
	Currency      sql.NullString  `json:"currency"`
	PartitionId   sql.NullInt64   `json:"partitionId"`
	Remark        sql.NullString  `json:"remark"`
	OtherInfo     sql.NullString  `json:"otherInfo"`
	ModCtr        int64           `json:"modCtr"`
	Created       sql.NullTime    `json:"created"`
	Updated       sql.NullTime    `json:"updated"`
}

const getGlAccount = `-- name: GetGlAccount :one
SELECT 
Id, mr.UUId, Office_Id, COA_Id, Balance, Pending_Trn_Amt, 
Account_Type_Id, Currency, Partition_Id, Remark, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM GL_Account d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesAccount) GetGlAccount(ctx context.Context, id int64) (GlAccountInfo, error) {
	row := q.db.QueryRowContext(ctx, getGlAccount, id)
	var i GlAccountInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeId,
		&i.CoaId,
		&i.Balance,
		&i.PendingTrnAmt,
		&i.AccountTypeId,
		&i.Currency,
		&i.PartitionId,
		&i.Remark,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getGlAccountbyUuid = `-- name: GetGlAccountbyUuid :one
SELECT 
Id, mr.UUId, Office_Id, COA_Id, Balance, Pending_Trn_Amt, 
Account_Type_Id, Currency, Partition_Id, Remark, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM GL_Account d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccount) GetGlAccountbyUuid(ctx context.Context, uuid uuid.UUID) (GlAccountInfo, error) {
	row := q.db.QueryRowContext(ctx, getGlAccountbyUuid, uuid)
	var i GlAccountInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeId,
		&i.CoaId,
		&i.Balance,
		&i.PendingTrnAmt,
		&i.AccountTypeId,
		&i.Currency,
		&i.PartitionId,
		&i.Remark,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listGlAccount = `-- name: ListGlAccount:many
SELECT 
Id, mr.UUId, Office_Id, COA_Id, Balance, Pending_Trn_Amt, 
Account_Type_Id, Currency, Partition_Id, Remark, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM GL_Account d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Office_Id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListGlAccountParams struct {
	OfficeId int64 `json:"officeId"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *QueriesAccount) ListGlAccount(ctx context.Context, arg ListGlAccountParams) ([]GlAccountInfo, error) {
	rows, err := q.db.QueryContext(ctx, listGlAccount, arg.OfficeId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GlAccountInfo{}
	for rows.Next() {
		var i GlAccountInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.OfficeId,
			&i.CoaId,
			&i.Balance,
			&i.PendingTrnAmt,
			&i.AccountTypeId,
			&i.Currency,
			&i.PartitionId,
			&i.Remark,
			&i.OtherInfo,

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

const updateGlAccount = `-- name: UpdateGlAccount :one
UPDATE GL_Account SET 
Office_Id = $2,
COA_Id = $3,
Balance = $4,
Pending_Trn_Amt = $5,
Account_Type_Id = $6,
Currency = $7,
Partition_Id = $8,
Remark = $9,
Other_Info = $10
WHERE id = $1
RETURNING Id, UUId, Office_Id, COA_Id, Balance, Pending_Trn_Amt, Account_Type_Id, 
Currency, Partition_Id, Remark, Other_Info
`

func (q *QueriesAccount) UpdateGlAccount(ctx context.Context, arg GlAccountRequest) (model.GlAccount, error) {
	row := q.db.QueryRowContext(ctx, updateGlAccount,
		arg.Id,
		arg.OfficeId,
		arg.CoaId,
		arg.Balance,
		arg.PendingTrnAmt,
		arg.AccountTypeId,
		arg.Currency,
		arg.PartitionId,
		arg.Remark,
		arg.OtherInfo,
	)
	var i model.GlAccount
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeId,
		&i.CoaId,
		&i.Balance,
		&i.PendingTrnAmt,
		&i.AccountTypeId,
		&i.Currency,
		&i.PartitionId,
		&i.Remark,
		&i.OtherInfo,
	)
	return i, err
}
