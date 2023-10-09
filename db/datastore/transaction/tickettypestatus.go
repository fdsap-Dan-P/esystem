package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createTicketTypeStatus = `-- name: CreateTicketTypeStatus: one
INSERT INTO Ticket_Type_Status(
   uuid, product_ticket_type_id, status_id, ticket_type_action_array, other_info )
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (uuid)
DO UPDATE SET 
	product_ticket_type_id =  EXCLUDED.product_ticket_type_id,
	status_id =  EXCLUDED.status_id,
	ticket_type_action_array =  EXCLUDED.ticket_type_action_array,
	other_info =  EXCLUDED.other_info
RETURNING id, uuid, product_ticket_type_id, status_id, ticket_type_action_array, other_info`

type TicketTypeStatusRequest struct {
	Id                    int64          `json:"id"`
	Uuid                  uuid.UUID      `json:"uuid"`
	ProductTicketTypeId   int64          `json:"productTicketTypeId"`
	StatusId              int64          `json:"statusId"`
	TicketTypeActionArray []int64        `json:"ticketTypeActionArray"`
	OtherInfo             sql.NullString `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateTicketTypeStatus(ctx context.Context, arg TicketTypeStatusRequest) (model.TicketTypeStatus, error) {
	row := q.db.QueryRowContext(ctx, createTicketTypeStatus,
		arg.Uuid,
		arg.ProductTicketTypeId,
		arg.StatusId,
		pq.Array(arg.TicketTypeActionArray),
		arg.OtherInfo,
	)
	var i model.TicketTypeStatus
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductTicketTypeId,
		&i.StatusId,
		pq.Array(&i.TicketTypeActionArray),
		&i.OtherInfo,
	)
	return i, err
}

const deleteTicketTypeStatus = `-- name: DeleteTicketTypeStatus :exec
DELETE FROM Ticket_Type_Status
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketTypeStatus(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketTypeStatus, uuid)
	return err
}

type TicketTypeStatusInfo struct {
	Id                    int64          `json:"id"`
	Uuid                  uuid.UUID      `json:"uuid"`
	ProductTicketTypeId   int64          `json:"productTicketTypeId"`
	StatusId              int64          `json:"statusId"`
	TicketTypeActionArray []int64        `json:"ticketTypeActionArray"`
	OtherInfo             sql.NullString `json:"otherInfo"`
	ModCtr                int64          `json:"modCtr"`
	Created               sql.NullTime   `json:"created"`
	Updated               sql.NullTime   `json:"updated"`
}

const ticketTypeStatusSQL = `-- name: TicketTypeStatusSQL :one
SELECT
 id, mr.uuid, product_ticket_type_id, status_id, ticket_type_action_array, other_info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Type_Status d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateTicketTypeStatus(q *QueriesTransaction, ctx context.Context, sql string) (TicketTypeStatusInfo, error) {
	var i TicketTypeStatusInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductTicketTypeId,
		&i.StatusId,
		pq.Array(&i.TicketTypeActionArray),
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketTypeStatuss(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketTypeStatusInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TicketTypeStatusInfo{}
	for rows.Next() {
		var i TicketTypeStatusInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.ProductTicketTypeId,
			&i.StatusId,
			pq.Array(&i.TicketTypeActionArray),
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

func (q *QueriesTransaction) GetTicketTypeStatus(ctx context.Context, id int64) (TicketTypeStatusInfo, error) {
	return populateTicketTypeStatus(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", ticketTypeStatusSQL, id))
}

func (q *QueriesTransaction) GetTicketTypeStatusbyUuid(ctx context.Context, uuid uuid.UUID) (TicketTypeStatusInfo, error) {
	return populateTicketTypeStatus(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", ticketTypeStatusSQL, uuid))
}

type ListTicketTypeStatusParams struct {
	ProductTicketTypeId int64 `json:"productTicketTypeId"`
	Limit               int32 `json:"limit"`
	Offset              int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketTypeStatus(ctx context.Context, arg ListTicketTypeStatusParams) ([]TicketTypeStatusInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			ticketTypeStatusSQL, arg.Limit, arg.Offset)
	} else {
		sql = ticketTypeStatusSQL
	}
	return populateTicketTypeStatuss(q, ctx, sql)
}

const updateTicketTypeStatus = `-- name: UpdateTicketTypeStatus :one
UPDATE Ticket_Type_Status SET 
	uuid = $2,
	product_ticket_type_id = $3,
	status_id = $4,
	ticket_type_action_array = $5,
	other_info = $6
WHERE id = $1
RETURNING id, uuid, product_ticket_type_id, status_id, ticket_type_action_array, other_info
`

func (q *QueriesTransaction) UpdateTicketTypeStatus(ctx context.Context, arg TicketTypeStatusRequest) (model.TicketTypeStatus, error) {
	row := q.db.QueryRowContext(ctx, updateTicketTypeStatus,
		arg.Id,
		arg.Uuid,
		arg.ProductTicketTypeId,
		arg.StatusId,
		pq.Array(arg.TicketTypeActionArray),
		arg.OtherInfo,
	)
	var i model.TicketTypeStatus
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductTicketTypeId,
		&i.StatusId,
		pq.Array(&i.TicketTypeActionArray),
		&i.OtherInfo,
	)
	return i, err
}
