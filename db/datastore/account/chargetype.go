package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createChargeType = `-- name: CreateChargeType: one
INSERT INTO Charge_Type (
Charge_Type, UnRealized_Id, Realized_Id, Other_Info
) VALUES (
$1, $2, $3, $4
) RETURNING Id, UUId, Charge_Type, UnRealized_Id, Realized_Id, Other_Info
`

type ChargeTypeRequest struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	ChargeType   string         `json:"chargeType"`
	UnrealizedId int64          `json:"unrealizedId"`
	RealizedId   int64          `json:"realizedId"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}

func (q *QueriesAccount) CreateChargeType(ctx context.Context, arg ChargeTypeRequest) (model.ChargeType, error) {
	row := q.db.QueryRowContext(ctx, createChargeType,
		arg.ChargeType,
		arg.UnrealizedId,
		arg.RealizedId,
		arg.OtherInfo,
	)
	var i model.ChargeType
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ChargeType,
		&i.UnrealizedId,
		&i.RealizedId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteChargeType = `-- name: DeleteChargeType :exec
DELETE FROM Charge_Type
WHERE id = $1
`

func (q *QueriesAccount) DeleteChargeType(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteChargeType, id)
	return err
}

type ChargeTypeInfo struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	ChargeType   string         `json:"chargeType"`
	UnrealizedId int64          `json:"unrealizedId"`
	RealizedId   int64          `json:"realizedId"`
	OtherInfo    sql.NullString `json:"otherInfo"`
	ModCtr       int64          `json:"modCtr"`
	Created      sql.NullTime   `json:"created"`
	Updated      sql.NullTime   `json:"updated"`
}

const getChargeType = `-- name: GetChargeType :one
SELECT 
Id, mr.UUId, Charge_Type, UnRealized_Id, Realized_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Charge_Type d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesAccount) GetChargeType(ctx context.Context, id int64) (ChargeTypeInfo, error) {
	row := q.db.QueryRowContext(ctx, getChargeType, id)
	var i ChargeTypeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ChargeType,
		&i.UnrealizedId,
		&i.RealizedId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getChargeTypebyUuid = `-- name: GetChargeTypebyUuid :one
SELECT 
Id, mr.UUId, Charge_Type, UnRealized_Id, Realized_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Charge_Type d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccount) GetChargeTypebyUuid(ctx context.Context, uuid uuid.UUID) (ChargeTypeInfo, error) {
	row := q.db.QueryRowContext(ctx, getChargeTypebyUuid, uuid)
	var i ChargeTypeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ChargeType,
		&i.UnrealizedId,
		&i.RealizedId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getChargeTypebyName = `-- name: GetChargeTypebyName :one
SELECT 
Id, mr.UUId, Charge_Type, UnRealized_Id, Realized_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Charge_Type d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Charge_Type = $1 LIMIT 1
`

func (q *QueriesAccount) GetChargeTypebyName(ctx context.Context, name string) (ChargeTypeInfo, error) {
	row := q.db.QueryRowContext(ctx, getChargeTypebyName, name)
	var i ChargeTypeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ChargeType,
		&i.UnrealizedId,
		&i.RealizedId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listChargeType = `-- name: ListChargeType:many
SELECT 
Id, mr.UUId, Charge_Type, UnRealized_Id, Realized_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Charge_Type d INNER JOIN Main_Record mr on mr.UUId = d.UUId
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListChargeTypeParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccount) ListChargeType(ctx context.Context, arg ListChargeTypeParams) ([]ChargeTypeInfo, error) {
	rows, err := q.db.QueryContext(ctx, listChargeType, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ChargeTypeInfo{}
	for rows.Next() {
		var i ChargeTypeInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.ChargeType,
			&i.UnrealizedId,
			&i.RealizedId,
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

const updateChargeType = `-- name: UpdateChargeType :one
UPDATE Charge_Type SET 
Charge_Type = $2,
UnRealized_Id = $3,
Realized_Id = $4,
Other_Info = $5
WHERE id = $1
RETURNING Id, UUId, Charge_Type, UnRealized_Id, Realized_Id, Other_Info
`

func (q *QueriesAccount) UpdateChargeType(ctx context.Context, arg ChargeTypeRequest) (model.ChargeType, error) {
	row := q.db.QueryRowContext(ctx, updateChargeType,
		arg.Id,
		arg.ChargeType,
		arg.UnrealizedId,
		arg.RealizedId,
		arg.OtherInfo,
	)
	var i model.ChargeType
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ChargeType,
		&i.UnrealizedId,
		&i.RealizedId,
		&i.OtherInfo,
	)
	return i, err
}
