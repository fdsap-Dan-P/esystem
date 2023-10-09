package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
)

const createSystemConfig = `-- name: CreateSystemConfig: one
INSERT INTO System_Config (
Office_Id, GL_Date, Last_Accruals, Last_Month_End, Next_Month_End, 
System_Date, Run_State, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8
) 
ON CONFLICT(Office_Id) DO UPDATE SET
	GL_Date = excluded.GL_Date,
	Last_Accruals = excluded.Last_Accruals,
	Last_Month_End = excluded.Last_Month_End,
	Next_Month_End = excluded.Next_Month_End,
	Run_State = excluded.Run_State,
	System_Date = excluded.System_Date,
	Other_Info = excluded.Other_Info
RETURNING UUId, Office_Id, GL_Date, Last_Accruals, Last_Month_End, Next_Month_End, 
System_Date, Run_State, Other_Info
`

type SystemConfigRequest struct {
	Uuid         uuid.UUID      `json:"uuid"`
	OfficeId     int64          `json:"officeId"`
	GlDate       time.Time      `json:"glDate"`
	LastAccruals time.Time      `json:"lastAccruals"`
	LastMonthEnd time.Time      `json:"lastMonthEnd"`
	NextMonthEnd time.Time      `json:"nextMonthEnd"`
	SystemDate   time.Time      `json:"systemDate"`
	RunState     int16          `json:"runState"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}

func (q *QueriesReference) CreateSystemConfig(ctx context.Context, arg SystemConfigRequest) (model.SystemConfig, error) {
	row := q.db.QueryRowContext(ctx, createSystemConfig,
		arg.OfficeId,
		arg.GlDate,
		arg.LastAccruals,
		arg.LastMonthEnd,
		arg.NextMonthEnd,
		arg.SystemDate,
		arg.RunState,
		arg.OtherInfo,
	)
	var i model.SystemConfig
	err := row.Scan(
		&i.Uuid,
		&i.OfficeId,
		&i.GlDate,
		&i.LastAccruals,
		&i.LastMonthEnd,
		&i.NextMonthEnd,
		&i.SystemDate,
		&i.RunState,
		&i.OtherInfo,
	)
	return i, err
}

const deleteSystemConfig = `-- name: DeleteSystemConfig :exec
DELETE FROM System_Config
WHERE uuid = $1
`

func (q *QueriesReference) DeleteSystemConfig(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSystemConfig, uuid)
	return err
}

type SystemConfigInfo struct {
	Uuid         uuid.UUID      `json:"uuid"`
	OfficeId     int64          `json:"officeId"`
	GlDate       time.Time      `json:"glDate"`
	LastAccruals time.Time      `json:"lastAccruals"`
	LastMonthEnd time.Time      `json:"lastMonthEnd"`
	NextMonthEnd time.Time      `json:"nextMonthEnd"`
	SystemDate   time.Time      `json:"systemDate"`
	RunState     int16          `json:"runState"`
	OtherInfo    sql.NullString `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getSystemConfig = `-- name: GetSystemConfig :one
SELECT 
mr.UUId, Office_Id, GL_Date, Last_Accruals, 
Last_Month_End, Next_Month_End, System_Date, Run_State, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM System_Config d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesReference) GetSystemConfig(ctx context.Context, uuid uuid.UUID) (SystemConfigInfo, error) {
	row := q.db.QueryRowContext(ctx, getSystemConfig, uuid)
	var i SystemConfigInfo
	err := row.Scan(
		&i.Uuid,
		&i.OfficeId,
		&i.GlDate,
		&i.LastAccruals,
		&i.LastMonthEnd,
		&i.NextMonthEnd,
		&i.SystemDate,
		&i.RunState,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getSystemConfigbyUuId = `-- name: GetSystemConfigbyUuId :one
SELECT 
mr.UUId, Office_Id, GL_Date, Last_Accruals, 
Last_Month_End, Next_Month_End, System_Date, Run_State, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM System_Config d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesReference) GetSystemConfigbyUuId(ctx context.Context, uuid uuid.UUID) (SystemConfigInfo, error) {
	row := q.db.QueryRowContext(ctx, getSystemConfigbyUuId, uuid)
	var i SystemConfigInfo
	err := row.Scan(
		&i.Uuid,
		&i.OfficeId,
		&i.GlDate,
		&i.LastAccruals,
		&i.LastMonthEnd,
		&i.NextMonthEnd,
		&i.SystemDate,
		&i.RunState,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listSystemConfig = `-- name: ListSystemConfig:many
SELECT 
mr.UUId, Office_Id, GL_Date, Last_Accruals, 
Last_Month_End, Next_Month_End, System_Date, Run_State, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM System_Config d INNER JOIN Main_Record mr on mr.UUId = d.UUId
ORDER BY uuid
LIMIT $1
OFFSET $2
`

type ListSystemConfigParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesReference) ListSystemConfig(ctx context.Context, arg ListSystemConfigParams) ([]SystemConfigInfo, error) {
	rows, err := q.db.QueryContext(ctx, listSystemConfig, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SystemConfigInfo{}
	for rows.Next() {
		var i SystemConfigInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.OfficeId,
			&i.GlDate,
			&i.LastAccruals,
			&i.LastMonthEnd,
			&i.NextMonthEnd,
			&i.SystemDate,
			&i.RunState,
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

const updateSystemConfig = `-- name: UpdateSystemConfig :one
UPDATE System_Config SET 
Office_Id = $2,
GL_Date = $3,
Last_Accruals = $4,
Last_Month_End = $5,
Next_Month_End = $6,
System_Date = $7,
Run_State = $8,
Other_Info = $9
WHERE uuid = $1
RETURNING UUId, Office_Id, GL_Date, Last_Accruals, Last_Month_End, Next_Month_End, 
System_Date, Run_State, Other_Info
`

func (q *QueriesReference) UpdateSystemConfig(ctx context.Context, arg SystemConfigRequest) (model.SystemConfig, error) {
	row := q.db.QueryRowContext(ctx, updateSystemConfig,

		arg.Uuid,
		arg.OfficeId,
		arg.GlDate,
		arg.LastAccruals,
		arg.LastMonthEnd,
		arg.NextMonthEnd,
		arg.SystemDate,
		arg.RunState,
		arg.OtherInfo,
	)
	var i model.SystemConfig
	err := row.Scan(
		&i.Uuid,
		&i.OfficeId,
		&i.GlDate,
		&i.LastAccruals,
		&i.LastMonthEnd,
		&i.NextMonthEnd,
		&i.SystemDate,
		&i.RunState,
		&i.OtherInfo,
	)
	return i, err
}
