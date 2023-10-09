package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const createaccessConfigNumber = `-- name: CreateAccessConfigNumber: one
INSERT INTO Access_Config_Number(
   role_id, config_id, value, value2, measure_id)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(role_id, config_id)
DO UPDATE SET
  value = Excluded.value,
  value2 = Excluded.value2
RETURNING UUID, role_id, config_code, config_id, value, value2, measure_id`

type AccessConfigNumberRequest struct {
	Uuid      uuid.UUID       `json:"uuid"`
	RoleId    int64           `json:"roleId"`
	ConfigId  int64           `json:"configId"`
	Value     decimal.Decimal `json:"value"`
	Value2    decimal.Decimal `json:"value2"`
	MeasureId sql.NullInt64   `json:"measureId"`
}

func (q *QueriesAccess) CreateAccessConfigNumber(ctx context.Context, arg AccessConfigNumberRequest) (model.AccessConfigNumber, error) {
	row := q.db.QueryRowContext(ctx, createaccessConfigNumber,
		arg.RoleId,
		arg.ConfigId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.AccessConfigNumber
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ConfigCode,
		&i.ConfigId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const deleteAccessConfigNumber = `-- name: DeleteAccessConfigNumber :exec
DELETE FROM Access_Config_Number
WHERE uuid = $1
`

func (q *QueriesAccess) DeleteAccessConfigNumber(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccessConfigNumber, uuid)
	return err
}

type AccessConfigNumberInfo struct {
	Uuid       uuid.UUID       `json:"uuid"`
	RoleId     int64           `json:"roleId"`
	ConfigCode int64           `json:"configCode"`
	ConfigId   int64           `json:"configId"`
	Value      decimal.Decimal `json:"value"`
	Value2     decimal.Decimal `json:"value2"`
	MeasureId  sql.NullInt64   `json:"measureId"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

func populateAccessConfigNumber(q *QueriesAccess, ctx context.Context, sql string) (AccessConfigNumberInfo, error) {
	var i AccessConfigNumberInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.RoleId,
		&i.ConfigCode,
		&i.ConfigId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccessConfigNumbers(q *QueriesAccess, ctx context.Context, sql string) ([]AccessConfigNumberInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccessConfigNumberInfo{}
	for rows.Next() {
		var i AccessConfigNumberInfo
		err := rows.Scan(
			&i.Uuid,
			&i.RoleId,
			&i.ConfigCode,
			&i.ConfigId,
			&i.Value,
			&i.Value2,
			&i.MeasureId,
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

const accessConfigNumberSQL = `-- name: AccessConfigNumberSQL :one
SELECT
mr.UUID, role_id, config_code, config_id, value, value2, measure_id
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Access_Config_Number d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func (q *QueriesAccess) GetAccessConfigNumber(ctx context.Context, roleId int64, typeID int64) (AccessConfigNumberInfo, error) {
	return populateAccessConfigNumber(q, ctx, fmt.Sprintf("%s WHERE d.role_id = %v and d.config_id = %v", accessConfigNumberSQL, roleId, typeID))
}

func (q *QueriesAccess) GetAccessConfigNumberbyUuid(ctx context.Context, uuid uuid.UUID) (AccessConfigNumberInfo, error) {
	return populateAccessConfigNumber(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", accessConfigNumberSQL, uuid))
}

type ListAccessConfigNumberParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesAccess) ListAccessConfigNumber(ctx context.Context, arg ListAccessConfigNumberParams) ([]AccessConfigNumberInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d",
			accessConfigNumberSQL, arg.Limit, arg.Offset)
	} else {
		sql = accessConfigNumberSQL
	}
	return populateAccessConfigNumbers(q, ctx, sql)
}

const updateAccessConfigNumber = `-- name: UpdateAccessConfigNumber :one
UPDATE Access_Config_Number SET 
role_id = $2,
config_id = $3,
value = $4,
value2 = $5
WHERE uuid = $1
RETURNING uuid, role_id, config_code, config_id, value, value2
`

func (q *QueriesAccess) UpdateAccessConfigNumber(ctx context.Context, arg AccessConfigNumberRequest) (model.AccessConfigNumber, error) {
	row := q.db.QueryRowContext(ctx, updateAccessConfigNumber,
		arg.Uuid,
		arg.RoleId,
		arg.ConfigId,
		arg.Value,
		arg.Value2,
	)
	var i model.AccessConfigNumber
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
