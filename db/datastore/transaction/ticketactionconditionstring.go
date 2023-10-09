package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createTicketActionConditionString = `-- name: CreateTicketActionConditionString: one
INSERT INTO Ticket_Action_Condition_String 
  (UUID, ticket_type_status_id, Item_ID, Condition_Id, Value) 
VALUES 
  ($1, $2, $3, $4, $5) 
ON CONFLICT( UUID ) DO UPDATE SET
  ticket_type_status_id = excluded.ticket_type_status_id,
  Item_ID = excluded.Item_ID,
  Condition_ID = excluded.Condition_ID,
  Value = excluded.Value
RETURNING 
  UUID, ticket_type_status_id, Item_ID, Condition_Id, Value
`

type TicketActionConditionStringRequest struct {
	Uuid               uuid.UUID `json:"uuid"`
	TicketTypeStatusId int64     `json:"TicketTypeStatusId"`
	ItemId             int64     `json:"itemid"`
	ConditionId        int64     `json:"conditionId"`
	Value              string    `json:"value"`
}

func (q *QueriesTransaction) CreateTicketActionConditionString(ctx context.Context, arg TicketActionConditionStringRequest) (model.TicketActionConditionString, error) {
	row := q.db.QueryRowContext(ctx, createTicketActionConditionString,
		arg.Uuid,
		arg.TicketTypeStatusId,
		arg.ItemId,
		arg.ConditionId,
		arg.Value,
	)
	var i model.TicketActionConditionString
	err := row.Scan(
		&i.Uuid,
		&i.TicketTypeStatusId,
		&i.ItemCode,
		&i.ItemId,
		&i.ConditionId,
		&i.Value,
	)
	return i, err
}

const updateTicketActionConditionString = `-- name: UpdateTicketActionConditionString :one
UPDATE Ticket_Action_Condition_String SET 
  ticket_type_status_id = $2,
  Item_ID = $3,
  Condition_ID = $4,
  Value = $5
WHERE uuid = $1
RETURNING UUID, ticket_type_status_id, Item_ID, Condition_ID, Value
`

func (q *QueriesTransaction) UpdateTicketActionConditionString(ctx context.Context, arg TicketActionConditionStringRequest) (model.TicketActionConditionString, error) {
	row := q.db.QueryRowContext(ctx, updateTicketActionConditionString,
		arg.Uuid,
		arg.TicketTypeStatusId,
		arg.ItemId,
		arg.ConditionId,
		arg.Value,
	)
	var i model.TicketActionConditionString
	err := row.Scan(
		&i.Uuid,
		&i.TicketTypeStatusId,
		&i.ItemCode,
		&i.ItemId,
		&i.ConditionId,
		&i.Value,
	)
	return i, err
}

const deleteTicketActionConditionString = `-- name: DeleteTicketActionConditionString :exec
DELETE FROM Ticket_Action_Condition_String
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketActionConditionString(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketActionConditionString, uuid)
	return err
}

type TicketActionConditionStringInfo struct {
	Uuid               uuid.UUID    `json:"uuid"`
	TicketTypeStatusId int64        `json:"TicketTypeStatusId"`
	ItemId             int64        `json:"itemid"`
	ItemCode           string       `json:"ItemCode"`
	Item               string       `json:"item"`
	ItemShortName      string       `json:"itemShortName"`
	ItemDescription    string       `json:"itemDescription"`
	ConditionId        int64        `json:"conditionId"`
	Condition          string       `json:"condition"`
	Value              string       `json:"value"`
	ModCtr             int64        `json:"mod_ctr"`
	Created            sql.NullTime `json:"created"`
	Updated            sql.NullTime `json:"updated"`
}

func populateTicketActionConditiontring(q *QueriesTransaction, ctx context.Context, sql string) (TicketActionConditionStringInfo, error) {
	var i TicketActionConditionStringInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.TicketTypeStatusId,
		&i.ItemCode,
		&i.ItemId,
		&i.Item,
		&i.ItemShortName,
		&i.ItemDescription,
		&i.ConditionId,
		&i.Condition,
		&i.Value,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketActionConditiontring2(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketActionConditionStringInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []TicketActionConditionStringInfo{}
	for rows.Next() {
		var i TicketActionConditionStringInfo
		err := rows.Scan(
			&i.Uuid,
			&i.TicketTypeStatusId,
			&i.ItemCode,
			&i.ItemId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
			&i.ConditionId,
			&i.Condition,
			&i.Value,

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

const TicketActionConditionStringSQL = `-- name: TicketActionConditionStringSQL
SELECT 
  mr.UUID, d.ticket_type_status_id, d.Item_Code, d.Item_ID, 
  i.Title Item, i.Short_Name Item_Short_Name, i.Remark Item_Description, 
  d.Condition_Id, c.Title Condition, d.Value,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Action_Condition_String d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference i on d.Item_Id =  i.Id
INNER JOIN Reference c on d.Condition_Id =  c.Id
`

func (q *QueriesTransaction) GetTicketActionConditionString(ctx context.Context, ticketTypeStatusId int64, itemId int64) (TicketActionConditionStringInfo, error) {
	return populateTicketActionConditiontring(q, ctx, fmt.Sprintf("%s WHERE d.Ticket_Type_Status_Id = %v and d.Item_Id = %v",
		TicketActionConditionStringSQL, ticketTypeStatusId, itemId))
}

func (q *QueriesTransaction) GetTicketActionConditionStringbyUuid(ctx context.Context, uuid uuid.UUID) (TicketActionConditionStringInfo, error) {
	return populateTicketActionConditiontring(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", TicketActionConditionStringSQL, uuid))
}

type ListTicketActionConditionStringParams struct {
	TicketTypeStatusId int64 `json:"TicketTypeStatusId"`
	Limit              int32 `json:"limit"`
	Offset             int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketActionConditionString(ctx context.Context, arg ListTicketActionConditionStringParams) ([]TicketActionConditionStringInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.ticket_type_status_id = %v LIMIT %d OFFSET %d",
			TicketActionConditionStringSQL, arg.TicketTypeStatusId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.ticket_type_status_id = %v ", TicketActionConditionStringSQL, arg.TicketTypeStatusId)
	}
	return populateTicketActionConditiontring2(q, ctx, sql)
}
