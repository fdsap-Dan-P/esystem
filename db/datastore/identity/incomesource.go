package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createIncomeSource = `-- name: CreateIncomeSource: one
INSERT INTO Income_Source (
IIId, Series, Source, Type_Id, Min_Income, 
Max_Income, Remarks, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8
) RETURNING UUId, IIId, Series, Source, Type_Id, Min_Income, 
Max_Income, Remarks, Other_Info
`

type IncomeSourceRequest struct {
	Uuid      uuid.UUID       `json:"uuid"`
	Iiid      int64           `json:"iiid"`
	Series    int16           `json:"series"`
	Source    string          `json:"source"`
	TypeId    int64           `json:"typeId"`
	MinIncome decimal.Decimal `json:"minIncome"`
	MaxIncome decimal.Decimal `json:"maxIncome"`
	Remarks   sql.NullString  `json:"remarks"`
	OtherInfo sql.NullString  `json:"otherInfo"`
}

func (q *QueriesIdentity) CreateIncomeSource(ctx context.Context, arg IncomeSourceRequest) (model.IncomeSource, error) {
	row := q.db.QueryRowContext(ctx, createIncomeSource,
		arg.Iiid,
		arg.Series,
		arg.Source,
		arg.TypeId,
		arg.MinIncome,
		arg.MaxIncome,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.IncomeSource
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Source,
		&i.TypeId,
		&i.MinIncome,
		&i.MaxIncome,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteIncomeSource = `-- name: DeleteIncomeSource :exec
DELETE FROM Income_Source
WHERE uuid = $1
`

func (q *QueriesIdentity) DeleteIncomeSource(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteIncomeSource, uuid)
	return err
}

type IncomeSourceInfo struct {
	Uuid      uuid.UUID       `json:"uuid"`
	Iiid      int64           `json:"iiid"`
	Series    int16           `json:"series"`
	Source    string          `json:"source"`
	TypeId    int64           `json:"typeId"`
	MinIncome decimal.Decimal `json:"minIncome"`
	MaxIncome decimal.Decimal `json:"maxIncome"`
	Remarks   sql.NullString  `json:"remarks"`
	OtherInfo sql.NullString  `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getIncomeSource = `-- name: GetIncomeSource :one
SELECT 
mr.UUId, IIId, Series, Source, 
Type_Id, Min_Income, Max_Income, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Income_Source d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesIdentity) GetIncomeSource(ctx context.Context, uuid uuid.UUID) (IncomeSourceInfo, error) {
	row := q.db.QueryRowContext(ctx, getIncomeSource, uuid)
	var i IncomeSourceInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Source,
		&i.TypeId,
		&i.MinIncome,
		&i.MaxIncome,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getIncomeSourcebyUuId = `-- name: GetIncomeSourcebyUuId :one
SELECT 
mr.UUId, IIId, Series, Source, 
Type_Id, Min_Income, Max_Income, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Income_Source d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesIdentity) GetIncomeSourcebyUuId(ctx context.Context, uuid uuid.UUID) (IncomeSourceInfo, error) {
	row := q.db.QueryRowContext(ctx, getIncomeSourcebyUuId, uuid)
	var i IncomeSourceInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Source,
		&i.TypeId,
		&i.MinIncome,
		&i.MaxIncome,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getIncomeSourcebyName = `-- name: GetIncomeSourcebyName :one
SELECT 
mr.UUId, IIId, Series, Source, 
Type_Id, Min_Income, Max_Income, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Income_Source d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Title = $1 LIMIT 1
`

func (q *QueriesIdentity) GetIncomeSourcebyName(ctx context.Context, name string) (IncomeSourceInfo, error) {
	row := q.db.QueryRowContext(ctx, getIncomeSourcebyName, name)
	var i IncomeSourceInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Source,
		&i.TypeId,
		&i.MinIncome,
		&i.MaxIncome,
		&i.Remarks,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listIncomeSource = `-- name: ListIncomeSource:many
SELECT 
mr.UUId, IIId, Series, Source, 
Type_Id, Min_Income, Max_Income, Remarks, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Income_Source d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Iiid = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListIncomeSourceParams struct {
	Iiid   int64 `json:"iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListIncomeSource(ctx context.Context, arg ListIncomeSourceParams) ([]IncomeSourceInfo, error) {
	rows, err := q.db.QueryContext(ctx, listIncomeSource, arg.Iiid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []IncomeSourceInfo{}
	for rows.Next() {
		var i IncomeSourceInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.Iiid,
			&i.Series,
			&i.Source,
			&i.TypeId,
			&i.MinIncome,
			&i.MaxIncome,
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

const updateIncomeSource = `-- name: UpdateIncomeSource :one
UPDATE Income_Source SET 
IIId = $2,
Series = $3,
Source = $4,
Type_Id = $5,
Min_Income = $6,
Max_Income = $7,
Remarks = $8,
Other_Info = $9
WHERE uuid = $1
RETURNING UUId, IIId, Series, Source, Type_Id, Min_Income, 
Max_Income, Remarks, Other_Info
`

func (q *QueriesIdentity) UpdateIncomeSource(ctx context.Context, arg IncomeSourceRequest) (model.IncomeSource, error) {
	row := q.db.QueryRowContext(ctx, updateIncomeSource,

		arg.Uuid,
		arg.Iiid,
		arg.Series,
		arg.Source,
		arg.TypeId,
		arg.MinIncome,
		arg.MaxIncome,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.IncomeSource
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.Source,
		&i.TypeId,
		&i.MinIncome,
		&i.MaxIncome,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
