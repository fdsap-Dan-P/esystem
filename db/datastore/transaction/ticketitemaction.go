package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"
	"time"

	"github.com/google/uuid"
)

const createTicketItemAction = `-- name: CreateTicketItemAction: one
INSERT INTO Ticket_Item_Action(
   uuid, ticket_item_id, trn_head_id, user_id, action_id, action_date, remarks, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, uuid, ticket_item_id, trn_head_id, user_id, action_id, action_date, remarks, other_info`

type TicketItemActionRequest struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	TicketItemId int64          `json:"ticketItemId"`
	TrnHeadId    int64          `json:"trnHeadId"`
	UserId       int64          `json:"userId"`
	ActionId     int64          `json:"actionId"`
	ActionDate   time.Time      `json:"actionDate"`
	Remarks      sql.NullString `json:"remarks"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateTicketItemAction(ctx context.Context, arg TicketItemActionRequest) (model.TicketItemAction, error) {
	row := q.db.QueryRowContext(ctx, createTicketItemAction,
		arg.Uuid,
		arg.TicketItemId,
		arg.TrnHeadId,
		arg.UserId,
		arg.ActionId,
		arg.ActionDate,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.TicketItemAction
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TicketItemId,
		&i.TrnHeadId,
		&i.UserId,
		&i.ActionId,
		&i.ActionDate,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteTicketItemAction = `-- name: DeleteTicketItemAction :exec
DELETE FROM Ticket_Item_Action
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketItemAction(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketItemAction, uuid)
	return err
}

type TicketItemActionInfo struct {
	Id           int64          `json:"id"`
	Uuid         uuid.UUID      `json:"uuid"`
	TicketItemId int64          `json:"ticketItemId"`
	TrnHeadId    int64          `json:"trnHeadId"`
	UserId       int64          `json:"userId"`
	ActionId     int64          `json:"actionId"`
	ActionDate   time.Time      `json:"actionDate"`
	Remarks      sql.NullString `json:"remarks"`
	OtherInfo    sql.NullString `json:"otherInfo"`
	ModCtr       int64          `json:"modCtr"`
	Created      sql.NullTime   `json:"created"`
	Updated      sql.NullTime   `json:"updated"`
}

const ticketItemActionSQL = `-- name: TicketItemActionSQL :one
SELECT
  id, mr.UUID, ticket_item_id, trn_head_id, user_id, action_id, action_date, remarks, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Item_Action d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateTicketItemAction(q *QueriesTransaction, ctx context.Context, sql string) (TicketItemActionInfo, error) {
	var i TicketItemActionInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TicketItemId,
		&i.TrnHeadId,
		&i.UserId,
		&i.ActionId,
		&i.ActionDate,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketItemActions(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketItemActionInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TicketItemActionInfo{}
	for rows.Next() {
		var i TicketItemActionInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.TicketItemId,
			&i.TrnHeadId,
			&i.UserId,
			&i.ActionId,
			&i.ActionDate,
			&i.Remarks,
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

func (q *QueriesTransaction) GetTicketItemAction(ctx context.Context, id int64) (TicketItemActionInfo, error) {
	return populateTicketItemAction(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", ticketItemActionSQL, id))
}

func (q *QueriesTransaction) GetTicketItemActionbyUuid(ctx context.Context, uuid uuid.UUID) (TicketItemActionInfo, error) {
	return populateTicketItemAction(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", ticketItemActionSQL, uuid))
}

type ListTicketItemActionParams struct {
	TicketItemId int64 `json:"ticketItemId"`
	Limit        int32 `json:"limit"`
	Offset       int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketItemAction(ctx context.Context, arg ListTicketItemActionParams) ([]TicketItemActionInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			ticketItemActionSQL, arg.Limit, arg.Offset)
	} else {
		sql = ticketItemActionSQL
	}
	return populateTicketItemActions(q, ctx, sql)
}

const updateTicketItemAction = `-- name: UpdateTicketItemAction :one
UPDATE Ticket_Item_Action SET 
	uuid = $2,
	ticket_item_id = $3,
	trn_head_id = $4,
	user_id = $5,
	action_id = $6,
	action_date = $7,
	remarks = $8,
	other_info = $9
WHERE id = $1
RETURNING id, uuid, ticket_item_id, trn_head_id, user_id, action_id, action_date, remarks, other_info
`

func (q *QueriesTransaction) UpdateTicketItemAction(ctx context.Context, arg TicketItemActionRequest) (model.TicketItemAction, error) {
	row := q.db.QueryRowContext(ctx, updateTicketItemAction,
		arg.Id,
		arg.Uuid,
		arg.TicketItemId,
		arg.TrnHeadId,
		arg.UserId,
		arg.ActionId,
		arg.ActionDate,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.TicketItemAction
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TicketItemId,
		&i.TrnHeadId,
		&i.UserId,
		&i.ActionId,
		&i.ActionDate,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
