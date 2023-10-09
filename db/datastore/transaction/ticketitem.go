package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createTicketItem = `-- name: CreateTicketItem: one
INSERT INTO Ticket_Item(
   uuid, ticket_id, item_id, item, status_id, remarks, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT(UUID)
DO UPDATE SET
	ticket_id =  EXCLUDED.ticket_id,
	item_id =  EXCLUDED.item_id,
	item =  EXCLUDED.item,
	status_id =  EXCLUDED.status_id,
	remarks =  EXCLUDED.remarks,
	other_info =  EXCLUDED.other_info
RETURNING id, uuid, ticket_id, item_id, item, status_id, remarks, other_info`

type TicketItemRequest struct {
	Id        int64          `json:"id"`
	Uuid      uuid.UUID      `json:"uuid"`
	TicketId  int64          `json:"ticketId"`
	ItemId    int64          `json:"itemid"`
	ItemCode  int64          `json:"ItemCode"`
	Item      string         `json:"item"`
	StatusId  int64          `json:"statusId"`
	Remarks   string         `json:"remarks"`
	OtherInfo sql.NullString `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateTicketItem(ctx context.Context, arg TicketItemRequest) (model.TicketItem, error) {
	row := q.db.QueryRowContext(ctx, createTicketItem,
		arg.Uuid,
		arg.TicketId,
		arg.ItemId,
		arg.Item,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.TicketItem
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TicketId,
		&i.ItemCode,
		&i.ItemId,
		&i.Item,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}

const deleteTicketItem = `-- name: DeleteTicketItem :exec
DELETE FROM Ticket_Item
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketItem(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketItem, uuid)
	return err
}

type TicketItemInfo struct {
	Id        int64          `json:"id"`
	Uuid      uuid.UUID      `json:"uuid"`
	TicketId  int64          `json:"ticketId"`
	ItemId    int64          `json:"itemid"`
	ItemCode  int64          `json:"ItemCode"`
	Item      string         `json:"item"`
	StatusId  int64          `json:"statusId"`
	Remarks   string         `json:"remarks"`
	OtherInfo sql.NullString `json:"otherInfo"`
	ModCtr    int64          `json:"modCtr"`
	Created   sql.NullTime   `json:"created"`
	Updated   sql.NullTime   `json:"updated"`
}

const ticketItemSQL = `-- name: TicketItemSQL :one
SELECT
  id, mr.UUID, ticket_id, item_id, item, status_id, remarks, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Item d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateTicketItem(q *QueriesTransaction, ctx context.Context, sql string) (TicketItemInfo, error) {
	var i TicketItemInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TicketId,
		&i.ItemCode,
		&i.ItemId,
		&i.Item,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketItems(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketItemInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TicketItemInfo{}
	for rows.Next() {
		var i TicketItemInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.TicketId,
			&i.ItemCode,
			&i.ItemId,
			&i.Item,
			&i.StatusId,
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

func (q *QueriesTransaction) GetTicketItem(ctx context.Context, id int64) (TicketItemInfo, error) {
	return populateTicketItem(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", ticketItemSQL, id))
}

func (q *QueriesTransaction) GetTicketItembyUuid(ctx context.Context, uuid uuid.UUID) (TicketItemInfo, error) {
	return populateTicketItem(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", ticketItemSQL, uuid))
}

type ListTicketItemParams struct {
	TicketId int64 `json:"ticketId"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketItem(ctx context.Context, arg ListTicketItemParams) ([]TicketItemInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			ticketItemSQL, arg.Limit, arg.Offset)
	} else {
		sql = ticketItemSQL
	}
	return populateTicketItems(q, ctx, sql)
}

const updateTicketItem = `-- name: UpdateTicketItem :one
UPDATE Ticket_Item SET 
uuid = $2,
ticket_id = $3,
item_id = $4,
item = $5,
status_id = $6,
remarks = $7,
other_info = $8
WHERE id = $1
RETURNING id, uuid, ticket_id, item_id, item, status_id, remarks, other_info
`

func (q *QueriesTransaction) UpdateTicketItem(ctx context.Context, arg TicketItemRequest) (model.TicketItem, error) {
	row := q.db.QueryRowContext(ctx, updateTicketItem,
		arg.Id,
		arg.Uuid,
		arg.TicketId,
		arg.ItemId,
		arg.Item,
		arg.StatusId,
		arg.Remarks,
		arg.OtherInfo,
	)
	var i model.TicketItem
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.TicketId,
		&i.ItemCode,
		&i.ItemId,
		&i.Item,
		&i.StatusId,
		&i.Remarks,
		&i.OtherInfo,
	)
	return i, err
}
