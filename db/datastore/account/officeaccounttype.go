package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createOfficeAccountType = `-- name: CreateOfficeAccountType: one
INSERT INTO Office_Account_Type (
Office_Account_Type, COA_Id, Other_Info
) VALUES (
$1, $2, $3
) RETURNING Id, UUId, Office_Account_Type, COA_Id, Other_Info
`

type OfficeAccountTypeRequest struct {
	Id                int64          `json:"id"`
	Uuid              uuid.UUID      `json:"uuid"`
	OfficeAccountType string         `json:"officeAccountType"`
	CoaId             int64          `json:"coaId"`
	OtherInfo         sql.NullString `json:"otherInfo"`
}

func (q *QueriesAccount) CreateOfficeAccountType(ctx context.Context, arg OfficeAccountTypeRequest) (model.OfficeAccountType, error) {
	row := q.db.QueryRowContext(ctx, createOfficeAccountType,
		arg.OfficeAccountType,
		arg.CoaId,
		arg.OtherInfo,
	)
	var i model.OfficeAccountType
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeAccountType,
		&i.CoaId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteOfficeAccountType = `-- name: DeleteOfficeAccountType :exec
DELETE FROM Office_Account_Type
WHERE id = $1
`

func (q *QueriesAccount) DeleteOfficeAccountType(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteOfficeAccountType, id)
	return err
}

type OfficeAccountTypeInfo struct {
	Id                int64          `json:"id"`
	Uuid              uuid.UUID      `json:"uuid"`
	OfficeAccountType string         `json:"officeAccountType"`
	CoaId             int64          `json:"coaId"`
	OtherInfo         sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getOfficeAccountType = `-- name: GetOfficeAccountType :one
SELECT 
Id, mr.UUId, Office_Account_Type, COA_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office_Account_Type d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE id = $1 LIMIT 1
`

func (q *QueriesAccount) GetOfficeAccountType(ctx context.Context, id int64) (OfficeAccountTypeInfo, error) {
	row := q.db.QueryRowContext(ctx, getOfficeAccountType, id)
	var i OfficeAccountTypeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeAccountType,
		&i.CoaId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getOfficeAccountTypebyName = `-- name: GetOfficeAccountTypebyName :one
SELECT 
Id, mr.UUId, Office_Account_Type, COA_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office_Account_Type d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE lower(Office_Account_Type) = lower($1) LIMIT 1
`

func (q *QueriesAccount) GetOfficeAccountTypebyName(ctx context.Context, name string) (OfficeAccountTypeInfo, error) {
	row := q.db.QueryRowContext(ctx, getOfficeAccountTypebyName, name)
	var i OfficeAccountTypeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeAccountType,
		&i.CoaId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getOfficeAccountTypebyUuid = `-- name: GetOfficeAccountTypebyUuid :one
SELECT 
Id, mr.UUId, Office_Account_Type, COA_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office_Account_Type d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccount) GetOfficeAccountTypebyUuid(ctx context.Context, uuid uuid.UUID) (OfficeAccountTypeInfo, error) {
	row := q.db.QueryRowContext(ctx, getOfficeAccountTypebyUuid, uuid)
	var i OfficeAccountTypeInfo
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeAccountType,
		&i.CoaId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listOfficeAccountType = `-- name: ListOfficeAccountType:many
SELECT 
Id, mr.UUId, Office_Account_Type, COA_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Office_Account_Type d INNER JOIN Main_Record mr on mr.UUId = d.UUId
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListOfficeAccountTypeParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccount) ListOfficeAccountType(ctx context.Context, arg ListOfficeAccountTypeParams) ([]OfficeAccountTypeInfo, error) {
	rows, err := q.db.QueryContext(ctx, listOfficeAccountType, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OfficeAccountTypeInfo{}
	for rows.Next() {
		var i OfficeAccountTypeInfo
		if err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.OfficeAccountType,
			&i.CoaId,
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

const updateOfficeAccountType = `-- name: UpdateOfficeAccountType :one
UPDATE Office_Account_Type SET 
Office_Account_Type = $2,
COA_Id = $3,
Other_Info = $4
WHERE id = $1
RETURNING Id, UUId, Office_Account_Type, COA_Id, Other_Info
`

func (q *QueriesAccount) UpdateOfficeAccountType(ctx context.Context, arg OfficeAccountTypeRequest) (model.OfficeAccountType, error) {
	row := q.db.QueryRowContext(ctx, updateOfficeAccountType,
		arg.Id,
		arg.OfficeAccountType,
		arg.CoaId,
		arg.OtherInfo,
	)
	var i model.OfficeAccountType
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.OfficeAccountType,
		&i.CoaId,
		&i.OtherInfo,
	)
	return i, err
}
