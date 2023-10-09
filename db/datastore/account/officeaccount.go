package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createOfficeAccount = `-- name: CreateOfficeAccount: one
INSERT INTO Office_Account (
  Office_Id, Type_Id, Currency, Partition_Id, Balance, 
  Pending_Trn_Amt, Budget, Last_Activity_Date, Status_Id, Remarks, Other_Info) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
ON CONFLICT(Office_Id , Type_Id, lower(Currency), COALESCE(Partition_Id ,0)) 
DO UPDATE SET
	Office_Id  = excluded.Office_Id ,
	Type_Id = excluded.Type_Id,
	Currency = excluded.Currency,
	Balance = excluded.Balance,
	Pending_Trn_Amt = excluded.Pending_Trn_Amt,
	Budget = excluded.Budget,
	Last_Activity_Date = excluded.Last_Activity_Date,
	Partition_Id  = excluded.Partition_Id,
	Status_Id  = excluded.Status_Id,
	Remarks = excluded.Remarks
RETURNING 
  Id, UUId, Office_Id, Type_Id, Currency, Partition_Id, Balance, 
  Pending_Trn_Amt, Budget, Last_Activity_Date, Status_Id, Remarks, Other_Info
`

type OfficeAccountRequest struct {
	Id               int64           `json:"id"`
	Uuid             uuid.UUID       `json:"uuid"`
	OfficeId         int64           `json:"officeId"`
	TypeId           int64           `json:"typeId"`
	Currency         string          `json:"currency"`
	PartitionId      sql.NullInt64   `json:"partitionId"`
	Balance          decimal.Decimal `json:"balance"`
	PendingTrnAmt    decimal.Decimal `json:"pendingTrnAmt"`
	Budget           decimal.Decimal `json:"budget"`
	LastActivityDate sql.NullTime    `json:"lastActivityDate"`
	StatusId         int64           `json:"statusId"`
	Remarks          sql.NullString  `json:"remarks"`
	OtherInfo        sql.NullString  `json:"otherInfo"`
}

func (q *QueriesAccount) CreateOfficeAccount(ctx context.Context, arg OfficeAccountRequest) (model.OfficeAccount, error) {
	row := q.db.QueryRowContext(ctx, createOfficeAccount,
		arg.OfficeId,
		arg.TypeId,
		arg.Currency,
		arg.PartitionId,
		arg.Balance,
		arg.PendingTrnAmt,
		arg.Budget,
		arg.LastActivityDate,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.OfficeAccount
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeId,
		&i.TypeId,
		&i.Currency,
		&i.PartitionId,
		&i.Balance,
		&i.PendingTrnAmt,
		&i.Budget,
		&i.LastActivityDate,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteOfficeAccount = `-- name: DeleteOfficeAccount :exec
DELETE FROM Office_Account
WHERE id = $1
`

func (q *QueriesAccount) DeleteOfficeAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteOfficeAccount, id)
	return err
}

type OfficeAccountInfo struct {
	Id               int64           `json:"id"`
	Uuid             uuid.UUID       `json:"uuid"`
	OfficeId         int64           `json:"officeId"`
	TypeId           int64           `json:"typeId"`
	Currency         string          `json:"currency"`
	PartitionId      sql.NullInt64   `json:"partitionId"`
	Balance          decimal.Decimal `json:"balance"`
	PendingTrnAmt    decimal.Decimal `json:"pendingTrnAmt"`
	Budget           decimal.Decimal `json:"budget"`
	LastActivityDate sql.NullTime    `json:"lastActivityDate"`
	StatusId         int64           `json:"statusId"`
	Remarks          sql.NullString  `json:"remarks"`
	OtherInfo        sql.NullString  `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getOfficeAccount = `-- name: GetOfficeAccount :one
SELECT 
Id, mr.UUId, 
Office_Id, Type_Id, Currency, Partition_Id, Balance, Pending_Trn_Amt, 
Budget, Last_Activity_Date, Status_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office_Account d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesAccount) GetOfficeAccount(ctx context.Context, id int64) (OfficeAccountInfo, error) {
	row := q.db.QueryRowContext(ctx, getOfficeAccount, id)
	var i OfficeAccountInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeId,
		&i.TypeId,
		&i.Currency,
		&i.PartitionId,
		&i.Balance,
		&i.PendingTrnAmt,
		&i.Budget,
		&i.LastActivityDate,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getOfficeAccountbyCode = `-- name: GetOfficeAccountbyCode :one
SELECT 
Id, mr.UUId, 
Office_Id, Type_Id, Currency, Partition_Id, Balance, Pending_Trn_Amt, 
Budget, Last_Activity_Date, Status_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office_Account d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Office_Id = $1 and Type_Id = $2 and lower(Currency) = lower($3) and COALESCE(Partition_Id,0) = $4 
LIMIT 1
`

func (q *QueriesAccount) GetOfficeAccountbyCode(ctx context.Context,
	officeId int64, typeId int64, currency string, partId int64) (OfficeAccountInfo, error) {
	row := q.db.QueryRowContext(ctx, getOfficeAccountbyCode,
		officeId, typeId, currency, partId)
	var i OfficeAccountInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeId,
		&i.TypeId,
		&i.Currency,
		&i.PartitionId,
		&i.Balance,
		&i.PendingTrnAmt,
		&i.Budget,
		&i.LastActivityDate,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getOfficeAccountbyUuid = `-- name: GetOfficeAccountbyUuid :one
SELECT 
Id, mr.UUId, 
Office_Id, Type_Id, Currency, Partition_Id, Balance, Pending_Trn_Amt, 
Budget, Last_Activity_Date, Status_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office_Account d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccount) GetOfficeAccountbyUuid(ctx context.Context, uuid uuid.UUID) (OfficeAccountInfo, error) {
	row := q.db.QueryRowContext(ctx, getOfficeAccountbyUuid, uuid)
	var i OfficeAccountInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeId,
		&i.TypeId,
		&i.Currency,
		&i.PartitionId,
		&i.Balance,
		&i.PendingTrnAmt,
		&i.Budget,
		&i.LastActivityDate,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listOfficeAccount = `-- name: ListOfficeAccount:many
SELECT 
Id, mr.UUId, 
Office_Id, Type_Id, Currency, Partition_Id, Balance, Pending_Trn_Amt, 
Budget, Last_Activity_Date, Status_Id, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office_Account d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Office_Id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListOfficeAccountParams struct {
	OfficeId int64 `json:"officeId"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *QueriesAccount) ListOfficeAccount(ctx context.Context, arg ListOfficeAccountParams) ([]OfficeAccountInfo, error) {
	rows, err := q.db.QueryContext(ctx, listOfficeAccount, arg.OfficeId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OfficeAccountInfo{}
	for rows.Next() {
		var i OfficeAccountInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.OfficeId,
			&i.TypeId,
			&i.Currency,
			&i.PartitionId,
			&i.Balance,
			&i.PendingTrnAmt,
			&i.Budget,
			&i.LastActivityDate,
			&i.StatusId,
			&i.Remarks,
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

const updateOfficeAccount = `-- name: UpdateOfficeAccount :one
UPDATE Office_Account SET 
Office_Id = $2,
Type_Id = $3,
Currency = $4,
Partition_Id = $5,
Balance = $6,
Pending_Trn_Amt = $7,
Budget = $8,
Last_Activity_Date = $9,
Status_Id = $10,
Remarks = $11,
Other_Info = $12
WHERE id = $1
RETURNING Id, UUId, Office_Id, Type_Id, Currency, Partition_Id, Balance, 
Pending_Trn_Amt, Budget, Last_Activity_Date, Status_Id, Remarks, Other_Info
`

func (q *QueriesAccount) UpdateOfficeAccount(ctx context.Context, arg OfficeAccountRequest) (model.OfficeAccount, error) {
	row := q.db.QueryRowContext(ctx, updateOfficeAccount,
		arg.Id,
		arg.OfficeId,
		arg.TypeId,
		arg.Currency,
		arg.PartitionId,
		arg.Balance,
		arg.PendingTrnAmt,
		arg.Budget,
		arg.LastActivityDate,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.OfficeAccount
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeId,
		&i.TypeId,
		&i.Currency,
		&i.PartitionId,
		&i.Balance,
		&i.PendingTrnAmt,
		&i.Budget,
		&i.LastActivityDate,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
