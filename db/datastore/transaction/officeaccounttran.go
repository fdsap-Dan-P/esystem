package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createOfficeAccountTran = `-- name: CreateOfficeAccountTran: one
INSERT INTO Office_Account_Tran (
Trn_Head_Id, Series, Office_Account_Id, Trn_Amt, Other_Info
) VALUES (
$1, $2, $3, $4, $5
) RETURNING UUId, Trn_Head_Id, Series, Office_Account_Id, Trn_Amt, Other_Info
`

type OfficeAccountTranRequest struct {
	Uuid            uuid.UUID       `json:"uuid"`
	TrnHeadId       int64           `json:"trnHeadId"`
	Series          int16           `json:"series"`
	OfficeAccountId int64           `json:"officeAccountId"`
	TrnAmt          decimal.Decimal `json:"trnAmt"`
	OtherInfo       sql.NullString  `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateOfficeAccountTran(ctx context.Context, arg OfficeAccountTranRequest) (model.OfficeAccountTran, error) {
	row := q.db.QueryRowContext(ctx, createOfficeAccountTran,
		arg.TrnHeadId,
		arg.Series,
		arg.OfficeAccountId,
		arg.TrnAmt,
		arg.OtherInfo,
	)
	var i model.OfficeAccountTran
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.OfficeAccountId,
		&i.TrnAmt,
		&i.OtherInfo,
	)
	return i, err
}

const deleteOfficeAccountTran = `-- name: DeleteOfficeAccountTran :exec
DELETE FROM Office_Account_Tran
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteOfficeAccountTran(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteOfficeAccountTran, uuid)
	return err
}

type OfficeAccountTranInfo struct {
	Uuid            uuid.UUID       `json:"uuid"`
	TrnHeadId       int64           `json:"trnHeadId"`
	Series          int16           `json:"series"`
	OfficeAccountId int64           `json:"officeAccountId"`
	TrnAmt          decimal.Decimal `json:"trnAmt"`
	OtherInfo       sql.NullString  `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getOfficeAccountTran = `-- name: GetOfficeAccountTran :one
SELECT 
mr.UUId, 
Trn_Head_Id, Series, Office_Account_Id, Trn_Amt, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office_Account_Tran d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesTransaction) GetOfficeAccountTran(ctx context.Context, uuid uuid.UUID) (OfficeAccountTranInfo, error) {
	row := q.db.QueryRowContext(ctx, getOfficeAccountTran, uuid)
	var i OfficeAccountTranInfo
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.OfficeAccountId,
		&i.TrnAmt,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getOfficeAccountTranbyUuid = `-- name: GetOfficeAccountTranbyUuid :one
SELECT 
mr.UUId, 
Trn_Head_Id, Series, Office_Account_Id, Trn_Amt, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office_Account_Tran d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesTransaction) GetOfficeAccountTranbyUuid(ctx context.Context, uuid uuid.UUID) (OfficeAccountTranInfo, error) {
	row := q.db.QueryRowContext(ctx, getOfficeAccountTranbyUuid, uuid)
	var i OfficeAccountTranInfo
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.OfficeAccountId,
		&i.TrnAmt,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listOfficeAccountTran = `-- name: ListOfficeAccountTran:many
SELECT 
mr.UUId, 
Trn_Head_Id, Series, Office_Account_Id, Trn_Amt, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office_Account_Tran d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE trn_head_id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListOfficeAccountTranParams struct {
	TrnHeadId int64 `json:"trnHeadId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesTransaction) ListOfficeAccountTran(ctx context.Context, arg ListOfficeAccountTranParams) ([]OfficeAccountTranInfo, error) {
	rows, err := q.db.QueryContext(ctx, listOfficeAccountTran, arg.TrnHeadId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OfficeAccountTranInfo{}
	for rows.Next() {
		var i OfficeAccountTranInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.TrnHeadId,
			&i.Series,
			&i.OfficeAccountId,
			&i.TrnAmt,
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

const updateOfficeAccountTran = `-- name: UpdateOfficeAccountTran :one
UPDATE Office_Account_Tran SET 
Trn_Head_Id = $2,
Series = $3,
Office_Account_Id = $4,
Trn_Amt = $5,
Other_Info = $6
WHERE uuid = $1
RETURNING UUId, Trn_Head_Id, Series, Office_Account_Id, Trn_Amt, Other_Info
`

func (q *QueriesTransaction) UpdateOfficeAccountTran(ctx context.Context, arg OfficeAccountTranRequest) (model.OfficeAccountTran, error) {
	row := q.db.QueryRowContext(ctx, updateOfficeAccountTran,

		arg.Uuid,
		arg.TrnHeadId,
		arg.Series,
		arg.OfficeAccountId,
		arg.TrnAmt,
		arg.OtherInfo,
	)
	var i model.OfficeAccountTran
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.OfficeAccountId,
		&i.TrnAmt,
		&i.OtherInfo,
	)
	return i, err
}
