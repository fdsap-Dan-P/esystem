package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createActionLink = `-- name: CreateActionLink: one
INSERT INTO Action_Link(
   uuid, event_name, type_id, end_point_call, server_id, other_info )
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(uuid)
DO UPDATE SET
	event_name =  EXCLUDED.event_name,
	type_id =  EXCLUDED.type_id,
	end_point_call =  EXCLUDED.end_point_call,
	server_id =  EXCLUDED.server_id,
	other_info =  EXCLUDED.other_info
RETURNING id, uuid, event_name, type_id, end_point_call, server_id, other_info`

type ActionLinkRequest struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	EventName    string         `json:"eventName"`
	TypeId       int64          `json:"typeId"`
	EndPointCall sql.NullString `json:"endPointCall"`
	ServerId     sql.NullInt64  `json:"serverId"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateActionLink(ctx context.Context, arg ActionLinkRequest) (model.ActionLink, error) {
	row := q.db.QueryRowContext(ctx, createActionLink,
		arg.Uuid,
		arg.EventName,
		arg.TypeId,
		arg.EndPointCall,
		arg.ServerId,
		arg.OtherInfo,
	)
	var i model.ActionLink
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.EventName,
		&i.TypeId,
		&i.EndPointCall,
		&i.ServerId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteActionLink = `-- name: DeleteActionLink :exec
DELETE FROM Action_Link
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteActionLink(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteActionLink, uuid)
	return err
}

type ActionLinkInfo struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	EventName    string         `json:"eventName"`
	TypeId       int64          `json:"typeId"`
	EndPointCall sql.NullString `json:"endPointCall"`
	ServerId     sql.NullInt64  `json:"serverId"`
	OtherInfo    sql.NullString `json:"otherInfo"`
	ModCtr       int64          `json:"modCtr"`
	Created      sql.NullTime   `json:"created"`
	Updated      sql.NullTime   `json:"updated"`
}

const actionLinkSQL = `-- name: ActionLinkSQL :one
SELECT
id, mr.UUID, event_name, type_id, end_point_call, server_id, other_info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Action_Link d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateActionLink(q *QueriesTransaction, ctx context.Context, sql string) (ActionLinkInfo, error) {
	var i ActionLinkInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.EventName,
		&i.TypeId,
		&i.EndPointCall,
		&i.ServerId,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateActionLinks(q *QueriesTransaction, ctx context.Context, sql string) ([]ActionLinkInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ActionLinkInfo{}
	for rows.Next() {
		var i ActionLinkInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.EventName,
			&i.TypeId,
			&i.EndPointCall,
			&i.ServerId,
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

func (q *QueriesTransaction) GetActionLink(ctx context.Context, id int64) (ActionLinkInfo, error) {
	return populateActionLink(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", actionLinkSQL, id))
}

func (q *QueriesTransaction) GetActionLinkbyUuid(ctx context.Context, uuid uuid.UUID) (ActionLinkInfo, error) {
	return populateActionLink(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", actionLinkSQL, uuid))
}

type ListActionLinkParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *QueriesTransaction) ListActionLink(ctx context.Context, arg ListActionLinkParams) ([]ActionLinkInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			actionLinkSQL, arg.Limit, arg.Offset)
	} else {
		sql = actionLinkSQL
	}
	return populateActionLinks(q, ctx, sql)
}

const updateActionLink = `-- name: UpdateActionLink :one
UPDATE Action_Link SET 
	uuid = $2,
	event_name = $3,
	type_id = $4,
	end_point_call = $5,
	server_id = $6,
	other_info = $7
WHERE id = $1
RETURNING id, uuid, event_name, type_id, end_point_call, server_id, other_info
`

func (q *QueriesTransaction) UpdateActionLink(ctx context.Context, arg ActionLinkRequest) (model.ActionLink, error) {
	row := q.db.QueryRowContext(ctx, updateActionLink,
		arg.Id,
		arg.Uuid,
		arg.EventName,
		arg.TypeId,
		arg.EndPointCall,
		arg.ServerId,
		arg.OtherInfo,
	)
	var i model.ActionLink
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.EventName,
		&i.TypeId,
		&i.EndPointCall,
		&i.ServerId,
		&i.OtherInfo,
	)
	return i, err
}
