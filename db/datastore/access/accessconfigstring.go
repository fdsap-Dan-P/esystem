package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createaccessConfigString = `-- name: CreateAccessConfigString: one
INSERT INTO Access_Config_String(
   role_id, config_id, value )
VALUES ($1, $2, $3)
ON CONFLICT(role_id, config_id)
DO UPDATE SET
  value = Excluded.value
RETURNING UUID, role_id, config_code, config_id, value`

type AccessConfigStringRequest struct {
	Uuid     uuid.UUID `json:"uuid"`
	RoleId   int64     `json:"roleId"`
	ConfigId int64     `json:"configId"`
	Value    string    `json:"value"`
}

func (q *QueriesAccess) CreateAccessConfigString(ctx context.Context, arg AccessConfigStringRequest) (model.AccessConfigString, error) {
	row := q.db.QueryRowContext(ctx, createaccessConfigString,
		arg.RoleId,
		arg.ConfigId,
		arg.Value,
	)
	var i model.AccessConfigString
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ConfigCode,
		&i.ConfigId,
		&i.Value,
	)
	return i, err
}

const deleteAccessConfigString = `-- name: DeleteAccessConfigString :exec
DELETE FROM Access_Config_String
WHERE uuid = $1
`

func (q *QueriesAccess) DeleteAccessConfigString(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccessConfigString, uuid)
	return err
}

type AccessConfigStringInfo struct {
	Uuid       uuid.UUID `json:"uuid"`
	RoleId     int64     `json:"roleId"`
	ConfigCode int64     `json:"configCode"`
	ConfigId   int64     `json:"configId"`
	Value      string    `json:"value"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

func populateAccessConfigString(q *QueriesAccess, ctx context.Context, sql string) (AccessConfigStringInfo, error) {
	var i AccessConfigStringInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ConfigCode,
		&i.ConfigId,
		&i.Value,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccessConfigStrings(q *QueriesAccess, ctx context.Context, sql string) ([]AccessConfigStringInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccessConfigStringInfo{}
	for rows.Next() {
		var i AccessConfigStringInfo
		err := rows.Scan(
			&i.Uuid,
			&i.RoleId,
			&i.ConfigCode,
			&i.ConfigId,
			&i.Value,
			&i.ModCtr,
			&i.Created,
			&i.Updated,
		)
		if err != nil {
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

const accessConfigStringSQL = `-- name: AccessConfigStringSQL :one
SELECT
mr.UUID, role_id, config_code, config_id, value
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Config_String d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func (q *QueriesAccess) GetAccessConfigString(ctx context.Context, roleId int64, typeID int64) (AccessConfigStringInfo, error) {
	return populateAccessConfigString(q, ctx, fmt.Sprintf("%s WHERE d.role_id = %v and d.config_id = %v", accessConfigStringSQL, roleId, typeID))
}

func (q *QueriesAccess) GetAccessConfigStringbyUuid(ctx context.Context, uuid uuid.UUID) (AccessConfigStringInfo, error) {
	return populateAccessConfigString(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", accessConfigStringSQL, uuid))
}

type ListAccessConfigStringParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccess) ListAccessConfigString(ctx context.Context, arg ListAccessConfigStringParams) ([]AccessConfigStringInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d",
			accessConfigStringSQL, arg.Limit, arg.Offset)
	} else {
		sql = accessConfigStringSQL
	}
	return populateAccessConfigStrings(q, ctx, sql)
}

const updateAccessConfigString = `-- name: UpdateAccessConfigString :one
UPDATE Access_Config_String SET 
role_id = $2,
config_id = $3,
value = $4
WHERE uuid = $1
RETURNING uuid, role_id, config_code, config_id, value
`

func (q *QueriesAccess) UpdateAccessConfigString(ctx context.Context, arg AccessConfigStringRequest) (model.AccessConfigString, error) {
	row := q.db.QueryRowContext(ctx, updateAccessConfigString,
		arg.Uuid,
		arg.RoleId,
		arg.ConfigId,
		arg.Value,
	)
	var i model.AccessConfigString
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ConfigCode,
		&i.ConfigId,
		&i.Value,
	)
	return i, err
}
