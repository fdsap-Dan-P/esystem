package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"
	"time"

	"github.com/google/uuid"
)

const createaccessConfigDate = `-- name: CreateAccessConfigDate: one
INSERT INTO Access_Config_Date(
   role_id, config_id, value, value2 )
VALUES ($1, $2, $3, $4)
ON CONFLICT(role_id, config_id)
DO UPDATE SET
  value = Excluded.value,
  value2 = Excluded.value2
RETURNING UUID, role_id, config_code, config_id, value, value2`

type AccessConfigDateRequest struct {
	Uuid     uuid.UUID `json:"uuid"`
	RoleId   int64     `json:"roleId"`
	ConfigId int64     `json:"configId"`
	Value    time.Time `json:"value"`
	Value2   time.Time `json:"value2"`
}

func (q *QueriesAccess) CreateAccessConfigDate(ctx context.Context, arg AccessConfigDateRequest) (model.AccessConfigDate, error) {
	row := q.db.QueryRowContext(ctx, createaccessConfigDate,
		arg.RoleId,
		arg.ConfigId,
		arg.Value,
		arg.Value2,
	)
	var i model.AccessConfigDate
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ConfigCode,
		&i.ConfigId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const deleteAccessConfigDate = `-- name: DeleteAccessConfigDate :exec
DELETE FROM Access_Config_Date
WHERE uuid = $1
`

func (q *QueriesAccess) DeleteAccessConfigDate(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccessConfigDate, uuid)
	return err
}

type AccessConfigDateInfo struct {
	Uuid       uuid.UUID `json:"uuid"`
	RoleId     int64     `json:"roleId"`
	ConfigCode int64     `json:"configCode"`
	ConfigId   int64     `json:"configId"`
	Value      time.Time `json:"value"`
	Value2     time.Time `json:"value2"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

func populateAccessConfigDate(q *QueriesAccess, ctx context.Context, sql string) (AccessConfigDateInfo, error) {
	var i AccessConfigDateInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ConfigCode,
		&i.ConfigId,
		&i.Value,
		&i.Value2,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccessConfigDates(q *QueriesAccess, ctx context.Context, sql string) ([]AccessConfigDateInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccessConfigDateInfo{}
	for rows.Next() {
		var i AccessConfigDateInfo
		err := rows.Scan(
			&i.Uuid,
			&i.RoleId,
			&i.ConfigCode,
			&i.ConfigId,
			&i.Value,
			&i.Value2,
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

const accessConfigDateSQL = `-- name: AccessConfigDateSQL :one
SELECT
mr.UUID, role_id, config_code, config_id, value, value2
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Config_Date d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func (q *QueriesAccess) GetAccessConfigDate(ctx context.Context, roleId int64, typeID int64) (AccessConfigDateInfo, error) {
	return populateAccessConfigDate(q, ctx, fmt.Sprintf("%s WHERE d.role_id = %v and d.config_id = %v", accessConfigDateSQL, roleId, typeID))
}

func (q *QueriesAccess) GetAccessConfigDatebyUuid(ctx context.Context, uuid uuid.UUID) (AccessConfigDateInfo, error) {
	return populateAccessConfigDate(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", accessConfigDateSQL, uuid))
}

type ListAccessConfigDateParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccess) ListAccessConfigDate(ctx context.Context, arg ListAccessConfigDateParams) ([]AccessConfigDateInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d",
			accessConfigDateSQL, arg.Limit, arg.Offset)
	} else {
		sql = accessConfigDateSQL
	}
	return populateAccessConfigDates(q, ctx, sql)
}

const updateAccessConfigDate = `-- name: UpdateAccessConfigDate :one
UPDATE Access_Config_Date SET 
role_id = $2,
config_id = $3,
value = $4,
value2 = $5
WHERE uuid = $1
RETURNING uuid, role_id, config_code, config_id, value, value2
`

func (q *QueriesAccess) UpdateAccessConfigDate(ctx context.Context, arg AccessConfigDateRequest) (model.AccessConfigDate, error) {
	row := q.db.QueryRowContext(ctx, updateAccessConfigDate,
		arg.Uuid,
		arg.RoleId,
		arg.ConfigId,
		arg.Value,
		arg.Value2,
	)
	var i model.AccessConfigDate
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ConfigCode,
		&i.ConfigId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}
