package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"
	"time"

	"github.com/google/uuid"
)

const createOfficer = `-- name: CreateOfficer: one
INSERT INTO Officer(
   uuid, office_id, position, period_start, period_end, status_id, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT(uuid)
DO UPDATE SET
	office_id =  EXCLUDED.office_id,
	position =  EXCLUDED.position,
	period_start =  EXCLUDED.period_start,
	period_end =  EXCLUDED.period_end,
	status_id =  EXCLUDED.status_id,
	other_info =  EXCLUDED.other_info
RETURNING UUID, office_id, position, period_start, period_end, status_id, other_info`

type OfficerRequest struct {
	Uuid        uuid.UUID      `json:"uuid"`
	OfficeId    int64          `json:"officeId"`
	Position    sql.NullString `json:"position"`
	PeriodStart time.Time      `json:"periodStart"`
	PeriodEnd   sql.NullTime   `json:"periodEnd"`
	StatusId    int64          `json:"statusId"`
	OtherInfo   sql.NullString `json:"otherInfo"`
}

func (q *QueriesIdentity) CreateOfficer(ctx context.Context, arg OfficerRequest) (model.Officer, error) {
	row := q.db.QueryRowContext(ctx, createOfficer,
		arg.Uuid,
		arg.OfficeId,
		arg.Position,
		arg.PeriodStart,
		arg.PeriodEnd,
		arg.StatusId,
		arg.OtherInfo,
	)
	var i model.Officer
	err := row.Scan(
		&i.Uuid,
		&i.OfficeId,
		&i.Position,
		&i.PeriodStart,
		&i.PeriodEnd,
		&i.StatusId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteOfficer = `-- name: DeleteOfficer :exec
DELETE FROM Officer
WHERE uuid = $1
`

func (q *QueriesIdentity) DeleteOfficer(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteOfficer, uuid)
	return err
}

type OfficerInfo struct {
	Uuid        uuid.UUID      `json:"uuid"`
	OfficeId    int64          `json:"officeId"`
	Position    sql.NullString `json:"position"`
	PeriodStart time.Time      `json:"periodStart"`
	PeriodEnd   sql.NullTime   `json:"periodEnd"`
	StatusId    int64          `json:"statusId"`
	OtherInfo   sql.NullString `json:"otherInfo"`
	ModCtr      int64          `json:"modCtr"`
	Created     sql.NullTime   `json:"created"`
	Updated     sql.NullTime   `json:"updated"`
}

const officerSQL = `-- name: OfficerSQL :one
SELECT
mr.UUID, office_id, position, period_start, period_end, status_id, other_info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Officer d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateOfficer(q *QueriesIdentity, ctx context.Context, sql string) (OfficerInfo, error) {
	var i OfficerInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.OfficeId,
		&i.Position,
		&i.PeriodStart,
		&i.PeriodEnd,
		&i.StatusId,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateOfficers(q *QueriesIdentity, ctx context.Context, sql string) ([]OfficerInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OfficerInfo{}
	for rows.Next() {
		var i OfficerInfo
		err := rows.Scan(
			&i.Uuid,
			&i.OfficeId,
			&i.Position,
			&i.PeriodStart,
			&i.PeriodEnd,
			&i.StatusId,
			&i.OtherInfo,
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

func (q *QueriesIdentity) GetOfficer(ctx context.Context, id int64) (OfficerInfo, error) {
	return populateOfficer(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", officerSQL, id))
}

func (q *QueriesIdentity) GetOfficerbyUuid(ctx context.Context, uuid uuid.UUID) (OfficerInfo, error) {
	return populateOfficer(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", officerSQL, uuid))
}

type ListOfficerParams struct {
	OfficeId int64 `json:"officeId"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *QueriesIdentity) ListOfficer(ctx context.Context, arg ListOfficerParams) ([]OfficerInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			officerSQL, arg.Limit, arg.Offset)
	} else {
		sql = officerSQL
	}
	return populateOfficers(q, ctx, sql)
}

const updateOfficer = `-- name: UpdateOfficer :one
UPDATE Officer SET 
	office_id = $2,
	position = $3,
	period_start = $4,
	period_end = $5,
	status_id = $6,
	other_info = $7
WHERE uuid = $1
RETURNING uuid, office_id, position, period_start, period_end, status_id, other_info
`

func (q *QueriesIdentity) UpdateOfficer(ctx context.Context, arg OfficerRequest) (model.Officer, error) {
	row := q.db.QueryRowContext(ctx, updateOfficer,
		arg.Uuid,
		arg.OfficeId,
		arg.Position,
		arg.PeriodStart,
		arg.PeriodEnd,
		arg.StatusId,
		arg.OtherInfo,
	)
	var i model.Officer
	err := row.Scan(
		&i.Uuid,
		&i.OfficeId,
		&i.Position,
		&i.PeriodStart,
		&i.PeriodEnd,
		&i.StatusId,
		&i.OtherInfo,
	)
	return i, err
}
