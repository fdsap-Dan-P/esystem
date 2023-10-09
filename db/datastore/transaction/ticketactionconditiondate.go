package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"
	"time"

	"github.com/google/uuid"
)

const createTicketActionConditionDate = `-- name: CreateTicketActionConditionDate: one
INSERT INTO Ticket_Action_Condition_Date(
   ticket_type_status_id, item_code, item_id, condition_id, value, value2 )
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (UUID)
DO UPDATE SET
  ticket_type_status_id = EXCLUDED.ticket_type_status_id,  
  item_id = EXCLUDED.item_id, 
  item_code = EXCLUDED.item_code,
  condition_id = EXCLUDED.condition_id, 
  value = EXCLUDED.value, 
  value2 = EXCLUDED.value2
RETURNING UUID, ticket_type_status_id, item_code, item_id, condition_id, value, value2`

type TicketActionConditionDateRequest struct {
	Uuid               uuid.UUID `json:"uuid"`
	TicketTypeStatusId int64     `json:"ticketTypeStatusId"`
	ItemId             int64     `json:"itemid"`
	ConditionId        int64     `json:"conditionId"`
	Value              time.Time `json:"value"`
	Value2             time.Time `json:"value2"`
}

func (q *QueriesTransaction) CreateTicketActionConditionDate(ctx context.Context, arg TicketActionConditionDateRequest) (model.TicketActionConditionDate, error) {
	row := q.db.QueryRowContext(ctx, createTicketActionConditionDate,
		arg.TicketTypeStatusId,
		arg.ItemId,
		arg.ConditionId,
		arg.Value,
		arg.Value2,
	)
	var i model.TicketActionConditionDate
	err := row.Scan(
		&i.Uuid,
		&i.TicketTypeStatusId,
		&i.ItemCode,
		&i.ItemId,
		&i.ConditionId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const deleteTicketActionConditionDate = `-- name: DeleteTicketActionConditionDate :exec
DELETE FROM Ticket_Action_Condition_Date
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketActionConditionDate(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketActionConditionDate, uuid)
	return err
}

type TicketActionConditionDateInfo struct {
	Uuid               uuid.UUID    `json:"uuid"`
	TicketTypeStatusId int64        `json:"ticketTypeStatusId"`
	ItemCode           string       `json:"itemCode"`
	ItemId             int64        `json:"itemid"`
	Item               string       `json:"item"`
	ConditionId        int64        `json:"conditionId"`
	Condition          string       `json:"condition"`
	Value              time.Time    `json:"value"`
	Value2             time.Time    `json:"value2"`
	ModCtr             int64        `json:"modCtr"`
	Created            sql.NullTime `json:"created"`
	Updated            sql.NullTime `json:"updated"`
}

const ticketActionConditionDateSQL = `-- name: TicketActionConditionDateSQL :one
SELECT
  mr.UUID, ticket_type_status_id, item_code, item_id, i.Title Item, condition_id, c.Short_Name Condition, value, value2
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Action_Condition_Date d 
INNER JOIN Reference i on d.Item_Id =  i.Id
INNER JOIN Reference c on d.Condition_Id =  c.Id
INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateTicketActionConditionDate(q *QueriesTransaction, ctx context.Context, sql string) (TicketActionConditionDateInfo, error) {
	var i TicketActionConditionDateInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.TicketTypeStatusId,
		&i.ItemCode,
		&i.ItemId,
		&i.Item,
		&i.ConditionId,
		&i.Condition,
		&i.Value,
		&i.Value2,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketActionConditionDates(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketActionConditionDateInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TicketActionConditionDateInfo{}
	for rows.Next() {
		var i TicketActionConditionDateInfo
		err := rows.Scan(
			&i.Uuid,
			&i.TicketTypeStatusId,
			&i.ItemCode,
			&i.ItemId,
			&i.Item,
			&i.ConditionId,
			&i.Condition,
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

func (q *QueriesTransaction) GetTicketActionConditionDate(ctx context.Context, ticketTypeStatusId int64, itemId int64) (TicketActionConditionDateInfo, error) {
	return populateTicketActionConditionDate(q, ctx, fmt.Sprintf("%s WHERE d.Ticket_Type_Status_Id = %v and d.Item_Id = %v", ticketActionConditionDateSQL, ticketTypeStatusId, itemId))
}

func (q *QueriesTransaction) GetTicketActionConditionDatebyUuid(ctx context.Context, uuid uuid.UUID) (TicketActionConditionDateInfo, error) {
	return populateTicketActionConditionDate(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", ticketActionConditionDateSQL, uuid))
}

type ListTicketActionConditionDateParams struct {
	TicketTypeStatusId int64 `json:"ticketTypeStatusId"`
	Limit              int32 `json:"limit"`
	Offset             int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketActionConditionDate(ctx context.Context, arg ListTicketActionConditionDateParams) ([]TicketActionConditionDateInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			ticketActionConditionDateSQL, arg.Limit, arg.Offset)
	} else {
		sql = ticketActionConditionDateSQL
	}
	return populateTicketActionConditionDates(q, ctx, sql)
}

const updateTicketActionConditionDate = `-- name: UpdateTicketActionConditionDate :one
UPDATE Ticket_Action_Condition_Date SET 
ticket_type_status_id = $2,
item_id = $4,
condition_id = $5,
value = $6,
value2 = $7
WHERE uuid = $1
RETURNING uuid, ticket_type_status_id, item_code, item_id, condition_id, value, value2
`

func (q *QueriesTransaction) UpdateTicketActionConditionDate(ctx context.Context, arg TicketActionConditionDateRequest) (model.TicketActionConditionDate, error) {
	row := q.db.QueryRowContext(ctx, updateTicketActionConditionDate,
		arg.Uuid,
		arg.TicketTypeStatusId,
		arg.ItemId,
		arg.ConditionId,
		arg.Value,
		arg.Value2,
	)
	var i model.TicketActionConditionDate
	err := row.Scan(
		&i.Uuid,
		&i.TicketTypeStatusId,
		&i.ItemCode,
		&i.ItemId,
		&i.ConditionId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}
