package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createIds = `-- name: CreateIds: one
INSERT INTO Ids (
IIId, Series, Id_Number, Registration_Date, Validity_Date, 
Type_Id, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7
) RETURNING UUId, IIId, Series, Id_Number, Registration_Date, Validity_Date, 
Type_Id, Other_Info
`

type IdsRequest struct {
	Uuid             uuid.UUID      `json:"uuid"`
	Iiid             int64          `json:"iiid"`
	Series           int16          `json:"series"`
	IdNumber         string         `json:"idNumber"`
	RegistrationDate sql.NullTime   `json:"registrationDate"`
	ValidityDate     sql.NullTime   `json:"ValidityDate"`
	TypeId           int64          `json:"typeId"`
	OtherInfo        sql.NullString `json:"otherInfo"`
}

func (q *QueriesIdentity) CreateIds(ctx context.Context, arg IdsRequest) (model.Ids, error) {
	row := q.db.QueryRowContext(ctx, createIds,
		arg.Iiid,
		arg.Series,
		arg.IdNumber,
		arg.RegistrationDate,
		arg.ValidityDate,
		arg.TypeId,
		arg.OtherInfo,
	)
	var i model.Ids
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.IdNumber,
		&i.RegistrationDate,
		&i.ValidityDate,
		&i.TypeId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteIds = `-- name: DeleteIds :exec
DELETE FROM Ids
WHERE uuid = $1
`

func (q *QueriesIdentity) DeleteIds(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteIds, uuid)
	return err
}

type IdsInfo struct {
	Uuid             uuid.UUID      `json:"uuid"`
	Iiid             int64          `json:"iiid"`
	Series           int16          `json:"series"`
	IdNumber         string         `json:"idNumber"`
	RegistrationDate sql.NullTime   `json:"registrationDate"`
	ValidityDate     sql.NullTime   `json:"ValidityDate"`
	TypeId           int64          `json:"typeId"`
	OtherInfo        sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getIds = `-- name: GetIds :one
SELECT 
  mr.UUId, IIId, Series, Id_Number, Registration_Date, 
	Validity_Date, Type_Id, Other_Info, mr.Mod_Ctr, 
	mr.Created, mr.Updated
FROM Ids d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesIdentity) GetIds(ctx context.Context, uuid uuid.UUID) (IdsInfo, error) {
	row := q.db.QueryRowContext(ctx, getIds, uuid)
	var i IdsInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.IdNumber,
		&i.RegistrationDate,
		&i.ValidityDate,
		&i.TypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getIdsbyUuId = `-- name: GetIdsbyUuId :one
SELECT 
mr.UUId, IIId, Series, 
Id_Number, Registration_Date, Validity_Date, Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ids d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesIdentity) GetIdsbyUuId(ctx context.Context, uuid uuid.UUID) (IdsInfo, error) {
	row := q.db.QueryRowContext(ctx, getIdsbyUuId, uuid)
	var i IdsInfo
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.IdNumber,
		&i.RegistrationDate,
		&i.ValidityDate,
		&i.TypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getbyIds = `-- name: GetbyIds :one
SELECT 
mr.UUId, IIId, Series, 
Id_Number, Registration_Date, Validity_Date, Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ids d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Id_Number LIKE '%' || $1 || '%' LIMIT $2 OFFSET $3
`

type GetbyIdsParams struct {
	IdNumber string `json:"IdNumber"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *QueriesIdentity) GetbyIds(ctx context.Context, arg GetbyIdsParams) ([]IdsInfo, error) {
	rows, err := q.db.QueryContext(ctx, getbyIds, arg.IdNumber, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []IdsInfo{}
	for rows.Next() {
		var i IdsInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.Iiid,
			&i.Series,
			&i.IdNumber,
			&i.RegistrationDate,
			&i.ValidityDate,
			&i.TypeId,
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

const listIds = `-- name: ListIds:many
SELECT 
mr.UUId, IIId, Series, 
Id_Number, Registration_Date, Validity_Date, Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ids d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Iiid = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListIdsParams struct {
	Iiid   int64 `json:"iiid"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesIdentity) ListIds(ctx context.Context, arg ListIdsParams) ([]IdsInfo, error) {
	rows, err := q.db.QueryContext(ctx, listIds, arg.Iiid, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []IdsInfo{}
	for rows.Next() {
		var i IdsInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.Iiid,
			&i.Series,
			&i.IdNumber,
			&i.RegistrationDate,
			&i.ValidityDate,
			&i.TypeId,
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

const updateIds = `-- name: UpdateIds :one
UPDATE Ids SET 
IIId = $2,
Series = $3,
Id_Number = $4,
Registration_Date = $5,
Validity_Date = $6,
Type_Id = $7,
Other_Info = $8
WHERE uuid = $1
RETURNING UUId, IIId, Series, Id_Number, Registration_Date, Validity_Date, 
Type_Id, Other_Info
`

func (q *QueriesIdentity) UpdateIds(ctx context.Context, arg IdsRequest) (model.Ids, error) {
	row := q.db.QueryRowContext(ctx, updateIds,

		arg.Uuid,
		arg.Iiid,
		arg.Series,
		arg.IdNumber,
		arg.RegistrationDate,
		arg.ValidityDate,
		arg.TypeId,
		arg.OtherInfo,
	)
	var i model.Ids
	err := row.Scan(
		&i.Uuid,
		&i.Iiid,
		&i.Series,
		&i.IdNumber,
		&i.RegistrationDate,
		&i.ValidityDate,
		&i.TypeId,
		&i.OtherInfo,
	)
	return i, err
}
