package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createUnitConversion = `-- name: CreateUnitConversion: one
INSERT INTO Unit_Conversion (
		Type_Id, From_Id, To_Id, Value, Other_Info)
	VALUES ($1, $2, $3, $4, $5)
	
	ON CONFLICT(From_ID, To_ID)
	DO UPDATE SET Type_ID = EXCLUDED.Type_ID, Value = EXCLUDED.Value
	
	RETURNING Id, UUID, Type_Id, From_Id, To_Id, Value, Other_Info 
`

// WITH
// val as
//  (SELECT $1::bigint Type_ID, $2::bigint From_Id, $3::bigint To_Id, $4::numeric(16,6) "value", $5::jsonb Other_Info),
// ins as
//  (INSERT INTO
//    Unit_Conversion(Type_ID, From_ID, To_ID, Value, Other_Info)

//    SELECT d.Type_Id, d.From_Id, d.To_Id, d.Value, d.Other_Info
//   FROM val d

// 	ON CONFLICT(From_Id, To_Id)
// 	DO UPDATE SET
// 	  Type_ID = EXCLUDED.Type_ID,
// 		Value = EXCLUDED.Value,
// 		Other_Info = EXCLUDED.Other_Info

// 	RETURNING From_Id, To_Id
//  )
// SELECT d.Id, d.UUID, d.Type_Id, d.From_Id, d.To_Id, d.Value, d.Other_Info
// FROM Unit_Conversion d
// LEFT JOIN ins v on d.From_ID = v.From_ID and d.To_ID = v.To_ID
//

// INSERT INTO Unit_Conversion (
// 	Type_Id, From_Id, To_Id, Value, Other_Info)
// VALUES ($1, $2, $3, $4, $5)

// ON CONFLICT(From_ID, To_ID)
// DO UPDATE SET Type_ID = EXCLUDED.Type_ID, Value = EXCLUDED.Value

// RETURNING Id, UUID, Type_Id, From_Id, To_Id, Value, Other_Info
type UnitConversionRequest struct {
	Id        int64           `json:"id"`
	Uuid      uuid.UUID       `json:"uuid"`
	TypeId    int64           `json:"typeId"`
	FromId    int64           `json:"fromId"`
	ToId      int64           `json:"toId"`
	Value     decimal.Decimal `json:"value"`
	OtherInfo sql.NullString  `json:"otherInfo"`
}

func (q *QueriesReference) CreateUnitConversion(ctx context.Context, arg UnitConversionRequest) (model.UnitConversion, error) {
	row := q.db.QueryRowContext(ctx, createUnitConversion,
		arg.TypeId,
		arg.FromId,
		arg.ToId,
		arg.Value,
		arg.OtherInfo,
	)
	var i model.UnitConversion
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TypeId,
		&i.FromId,
		&i.ToId,
		&i.Value,
		&i.OtherInfo,
	)
	fmt.Printf("ERROR createUnitConversion1: %+v\n", i.OtherInfo.String)

	fmt.Printf("to createUnitConversion1: %+v\n", i.FromId)
	fmt.Printf("to createUnitConversion2: %+v\n", arg.FromId)
	return i, err
}

const deleteUnitConversion = `-- name: DeleteUnitConversion :exec
DELETE FROM Unit_Conversion
WHERE id = $1
`

func (q *QueriesReference) DeleteUnitConversion(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUnitConversion, id)
	return err
}

type UnitConversionInfo struct {
	Id        int64           `json:"id"`
	Uuid      uuid.UUID       `json:"uuid"`
	TypeId    int64           `json:"typeId"`
	FromId    int64           `json:"fromId"`
	ToId      int64           `json:"toId"`
	Value     decimal.Decimal `json:"value"`
	OtherInfo sql.NullString  `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getUnitConversion = `-- name: GetUnitConversion :one
SELECT 
Id, mr.UUId, 
Type_Id, From_Id, To_Id, Value, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Unit_Conversion d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesReference) GetUnitConversion(ctx context.Context, id int64) (UnitConversionInfo, error) {
	row := q.db.QueryRowContext(ctx, getUnitConversion, id)
	var i UnitConversionInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TypeId,
		&i.FromId,
		&i.ToId,
		&i.Value,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUnitConversionbyAcc = `-- name: GetUnitConversionbyAcc :one
SELECT 
Id, mr.UUId, 
Type_Id, From_Id, To_Id, Value, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Unit_Conversion d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE acc = $1 LIMIT 1
`

func (q *QueriesReference) GetUnitConversionbyAcc(ctx context.Context, acc string) (UnitConversionInfo, error) {
	row := q.db.QueryRowContext(ctx, getUnitConversionbyAcc, acc)
	var i UnitConversionInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TypeId,
		&i.FromId,
		&i.ToId,
		&i.Value,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUnitConversionbyAltAcc = `-- name: GetUnitConversionbyAltAcc :one
SELECT 
Id, mr.UUId, 
Type_Id, From_Id, To_Id, Value, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Unit_Conversion d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE d.alternate_acc = $1 LIMIT 1
`

func (q *QueriesReference) GetUnitConversionbyAltAcc(ctx context.Context, altAcc string) (UnitConversionInfo, error) {
	row := q.db.QueryRowContext(ctx, getUnitConversionbyAltAcc, altAcc)
	var i UnitConversionInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TypeId,
		&i.FromId,
		&i.ToId,
		&i.Value,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getUnitConversionbyUuId = `-- name: GetUnitConversionbyUuId :one
SELECT 
Id, mr.UUId, 
Type_Id, From_Id, To_Id, Value, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Unit_Conversion d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesReference) GetUnitConversionbyUuId(ctx context.Context, uuid uuid.UUID) (UnitConversionInfo, error) {
	row := q.db.QueryRowContext(ctx, getUnitConversionbyUuId, uuid)
	var i UnitConversionInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TypeId,
		&i.FromId,
		&i.ToId,
		&i.Value,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listUnitConversion = `-- name: ListUnitConversion:many
SELECT 
Id, mr.UUId, 
Type_Id, From_Id, To_Id, Value, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Unit_Conversion d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Type_Id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListUnitConversionParams struct {
	TypeId int64 `json:"typeId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesReference) ListUnitConversion(ctx context.Context, arg ListUnitConversionParams) ([]UnitConversionInfo, error) {
	rows, err := q.db.QueryContext(ctx, listUnitConversion, arg.TypeId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UnitConversionInfo{}
	for rows.Next() {
		var i UnitConversionInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.TypeId,
			&i.FromId,
			&i.ToId,
			&i.Value,
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

const updateUnitConversion = `-- name: UpdateUnitConversion :one
UPDATE Unit_Conversion SET 
Type_Id = $2,
From_Id = $3,
To_Id = $4,
Value = $5,
Other_Info = $6
WHERE id = $1
RETURNING Id, UUId, Type_Id, From_Id, To_Id, Value, Other_Info
`

func (q *QueriesReference) UpdateUnitConversion(ctx context.Context, arg UnitConversionRequest) (model.UnitConversion, error) {
	row := q.db.QueryRowContext(ctx, updateUnitConversion,
		arg.Id,
		arg.TypeId,
		arg.FromId,
		arg.ToId,
		arg.Value,
		arg.OtherInfo,
	)
	var i model.UnitConversion
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TypeId,
		&i.FromId,
		&i.ToId,
		&i.Value,
		&i.OtherInfo,
	)
	return i, err
}
