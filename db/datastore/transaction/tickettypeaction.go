package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createTicketTypeAction = `-- name: CreateTicketTypeAction: one
INSERT INTO Ticket_Type_Action(
   uuid, product_ticket_type_id, action_id, actiondesc, action_link_id, other_info )
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT(UUID)
DO UPDATE SET
	product_ticket_type_id =  EXCLUDED.product_ticket_type_id,
	action_id =  EXCLUDED.action_id,
	actiondesc =  EXCLUDED.actiondesc,
	action_link_id =  EXCLUDED.action_link_id,
	other_info =  EXCLUDED.other_info
RETURNING id, uuid, product_ticket_type_id, action_id, actiondesc, action_link_id, other_info`

type TicketTypeActionRequest struct {
	Id                  int64          `json:"id"`
	Uuid                uuid.UUID      `json:"uuid"`
	ProductTicketTypeId int64          `json:"productTicketTypeId"`
	ActionId            int64          `json:"actionId"`
	Actiondesc          string         `json:"actiondesc"`
	ActionLinkId        sql.NullInt64  `json:"actionLinkId"`
	OtherInfo           sql.NullString `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateTicketTypeAction(ctx context.Context, arg TicketTypeActionRequest) (model.TicketTypeAction, error) {
	row := q.db.QueryRowContext(ctx, createTicketTypeAction,
		arg.Uuid,
		arg.ProductTicketTypeId,
		arg.ActionId,
		arg.Actiondesc,
		arg.ActionLinkId,
		arg.OtherInfo,
	)
	var i model.TicketTypeAction
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductTicketTypeId,
		&i.ActionId,
		&i.Actiondesc,
		&i.ActionLinkId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteTicketTypeAction = `-- name: DeleteTicketTypeAction :exec
DELETE FROM Ticket_Type_Action
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketTypeAction(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketTypeAction, uuid)
	return err
}

type TicketTypeActionInfo struct {
	Id                  int64          `json:"id"`
	Uuid                uuid.UUID      `json:"uuid"`
	ProductTicketTypeId int64          `json:"productTicketTypeId"`
	ActionId            int64          `json:"actionId"`
	Actiondesc          string         `json:"actiondesc"`
	ActionLinkId        sql.NullInt64  `json:"actionLinkId"`
	OtherInfo           sql.NullString `json:"otherInfo"`
	ModCtr              int64          `json:"modCtr"`
	Created             sql.NullTime   `json:"created"`
	Updated             sql.NullTime   `json:"updated"`
}

const ticketTypeActionSQL = `-- name: TicketTypeActionSQL :one
SELECT
	id, mr.UUID, product_ticket_type_id, action_id, actiondesc, action_link_id, other_info
	,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Type_Action d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateTicketTypeAction(q *QueriesTransaction, ctx context.Context, sql string) (TicketTypeActionInfo, error) {
	var i TicketTypeActionInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductTicketTypeId,
		&i.ActionId,
		&i.Actiondesc,
		&i.ActionLinkId,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketTypeActions(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketTypeActionInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TicketTypeActionInfo{}
	for rows.Next() {
		var i TicketTypeActionInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.ProductTicketTypeId,
			&i.ActionId,
			&i.Actiondesc,
			&i.ActionLinkId,
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

func (q *QueriesTransaction) GetTicketTypeAction(ctx context.Context, id int64) (TicketTypeActionInfo, error) {
	return populateTicketTypeAction(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", ticketTypeActionSQL, id))
}

func (q *QueriesTransaction) GetTicketTypeActionbyUuid(ctx context.Context, uuid uuid.UUID) (TicketTypeActionInfo, error) {
	return populateTicketTypeAction(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", ticketTypeActionSQL, uuid))
}

type ListTicketTypeActionParams struct {
	ProductTicketTypeId int64 `json:"ProductTicketTypeId"`
	Limit               int32 `json:"limit"`
	Offset              int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketTypeAction(ctx context.Context, arg ListTicketTypeActionParams) ([]TicketTypeActionInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			ticketTypeActionSQL, arg.Limit, arg.Offset)
	} else {
		sql = ticketTypeActionSQL
	}
	return populateTicketTypeActions(q, ctx, sql)
}

const updateTicketTypeAction = `-- name: UpdateTicketTypeAction :one
UPDATE Ticket_Type_Action SET 
uuid = $2,
product_ticket_type_id = $3,
action_id = $4,
actiondesc = $5,
action_link_id = $6,
other_info = $7
WHERE id = $1
RETURNING id, uuid, product_ticket_type_id, action_id, actiondesc, action_link_id, other_info
`

func (q *QueriesTransaction) UpdateTicketTypeAction(ctx context.Context, arg TicketTypeActionRequest) (model.TicketTypeAction, error) {
	row := q.db.QueryRowContext(ctx, updateTicketTypeAction,
		arg.Id,
		arg.Uuid,
		arg.ProductTicketTypeId,
		arg.ActionId,
		arg.Actiondesc,
		arg.ActionLinkId,
		arg.OtherInfo,
	)
	var i model.TicketTypeAction
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductTicketTypeId,
		&i.ActionId,
		&i.Actiondesc,
		&i.ActionLinkId,
		&i.OtherInfo,
	)
	return i, err
}
