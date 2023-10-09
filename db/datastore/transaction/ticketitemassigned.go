package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"
	"time"

	"github.com/google/uuid"
)

const createTicketItemAssigned = `-- name: CreateTicketItemAssigned: one
INSERT INTO Ticket_Item_Assigned(
   UUID, ticket_item_id, user_id, assigned_by_id, assigned_date, remarks, status_id, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT(UUID)
DO UPDATE SET
	ticket_item_id =  EXCLUDED.ticket_item_id,
	user_id =  EXCLUDED.user_id,
	assigned_by_id =  EXCLUDED.assigned_by_id,
	assigned_date =  EXCLUDED.assigned_date,
	remarks =  EXCLUDED.remarks,
	status_id =  EXCLUDED.status_id,
	other_info =  EXCLUDED.other_info
RETURNING UUID, ticket_item_id, user_id, assigned_by_id, assigned_date, remarks, status_id, other_info`

type TicketItemAssignedRequest struct {
	Uuid         uuid.UUID      `json:"uuid"`
	TicketItemId int64          `json:"ticketItemId"`
	UserId       int64          `json:"userId"`
	AssignedById int64          `json:"assignedById"`
	AssignedDate time.Time      `json:"assignedDate"`
	Remarks      string         `json:"remarks"`
	StatusId     int64          `json:"statusId"`
	OtherInfo    sql.NullString `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateTicketItemAssigned(ctx context.Context, arg TicketItemAssignedRequest) (model.TicketItemAssigned, error) {
	row := q.db.QueryRowContext(ctx, createTicketItemAssigned,
		arg.Uuid,
		arg.TicketItemId,
		arg.UserId,
		arg.AssignedById,
		arg.AssignedDate,
		arg.Remarks,
		arg.StatusId,
		arg.OtherInfo,
	)
	var i model.TicketItemAssigned
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.UserId,
		&i.AssignedById,
		&i.AssignedDate,
		&i.Remarks,
		&i.StatusId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteTicketItemAssigned = `-- name: DeleteTicketItemAssigned :exec
DELETE FROM Ticket_Item_Assigned
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketItemAssigned(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketItemAssigned, uuid)
	return err
}

type TicketItemAssignedInfo struct {
	Uuid         uuid.UUID      `json:"uuid"`
	TicketItemId int64          `json:"ticketItemId"`
	UserId       int64          `json:"userId"`
	AssignedById int64          `json:"assignedById"`
	AssignedDate time.Time      `json:"assignedDate"`
	Remarks      string         `json:"remarks"`
	StatusId     int64          `json:"statusId"`
	OtherInfo    sql.NullString `json:"otherInfo"`
	ModCtr       int64          `json:"modCtr"`
	Created      sql.NullTime   `json:"created"`
	Updated      sql.NullTime   `json:"updated"`
}

const ticketItemAssignedSQL = `-- name: TicketItemAssignedSQL :one
SELECT
mr.UUID, ticket_item_id, user_id, assigned_by_id, assigned_date, remarks, status_id, other_info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Item_Assigned d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateTicketItemAssigned(q *QueriesTransaction, ctx context.Context, sql string) (TicketItemAssignedInfo, error) {
	var i TicketItemAssignedInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.UserId,
		&i.AssignedById,
		&i.AssignedDate,
		&i.Remarks,
		&i.StatusId,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketItemAssigneds(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketItemAssignedInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TicketItemAssignedInfo{}
	for rows.Next() {
		var i TicketItemAssignedInfo
		err := rows.Scan(
			&i.Uuid,
			&i.TicketItemId,
			&i.UserId,
			&i.AssignedById,
			&i.AssignedDate,
			&i.Remarks,
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

// func (q *QueriesTransaction) GetTicketItemAssigned(ctx context.Context, id int64) (TicketItemAssignedInfo, error) {
// 	return populateTicketItemAssigned(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", ticketItemAssignedSQL, id))
// }

func (q *QueriesTransaction) GetTicketItemAssignedbyUuid(ctx context.Context, uuid uuid.UUID) (TicketItemAssignedInfo, error) {
	return populateTicketItemAssigned(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", ticketItemAssignedSQL, uuid))
}

type ListTicketItemAssignedParams struct {
	TicketItemId int64 `json:"ticketItemId"`
	Limit        int32 `json:"limit"`
	Offset       int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketItemAssigned(ctx context.Context, arg ListTicketItemAssignedParams) ([]TicketItemAssignedInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			ticketItemAssignedSQL, arg.Limit, arg.Offset)
	} else {
		sql = ticketItemAssignedSQL
	}
	return populateTicketItemAssigneds(q, ctx, sql)
}

const updateTicketItemAssigned = `-- name: UpdateTicketItemAssigned :one
UPDATE Ticket_Item_Assigned SET 
	ticket_item_id = $2,
	user_id = $3,
	assigned_by_id = $4,
	assigned_date = $5,
	remarks = $6,
	status_id = $7,
	other_info = $8
WHERE uuid = $1
RETURNING UUID, ticket_item_id, user_id, assigned_by_id, assigned_date, remarks, status_id, other_info
`

func (q *QueriesTransaction) UpdateTicketItemAssigned(ctx context.Context, arg TicketItemAssignedRequest) (model.TicketItemAssigned, error) {
	row := q.db.QueryRowContext(ctx, updateTicketItemAssigned,
		arg.Uuid,
		arg.TicketItemId,
		arg.UserId,
		arg.AssignedById,
		arg.AssignedDate,
		arg.Remarks,
		arg.StatusId,
		arg.OtherInfo,
	)
	var i model.TicketItemAssigned
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.UserId,
		&i.AssignedById,
		&i.AssignedDate,
		&i.Remarks,
		&i.StatusId,
		&i.OtherInfo,
	)
	return i, err
}
