package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createTicketActionConditionRef = `-- name: CreateTicketActionConditionRef: one
INSERT INTO Ticket_Action_Condition_Ref 
  (ticket_type_status_id, Item_ID, Condition_Id, Ref_Id) 
VALUES 
  ($1, $2, $3, $4) 
ON CONFLICT(UUID) DO UPDATE SET
  ticket_type_status_id = excluded.ticket_type_status_id,
  Item_Id = excluded.Item_ID,
  Condition_Id = excluded.Condition_Id,
  Ref_Id = excluded.Ref_Id
RETURNING 
  UUID, ticket_type_status_id, Item_Code, Item_ID, Condition_Id, Ref_Id
`

type TicketActionConditionRefRequest struct {
	Uuid               uuid.UUID `json:"uuid"`
	TicketTypeStatusId int64     `json:"ticketTypeStatusId"`
	ItemId             int64     `json:"itemid"`
	ConditionId        int64     `json:"conditionId"`
	RefId              int64     `json:"refId"`
}

func (q *QueriesTransaction) CreateTicketActionConditionRef(ctx context.Context, arg TicketActionConditionRefRequest) (model.TicketActionConditionRef, error) {
	row := q.db.QueryRowContext(ctx, createTicketActionConditionRef,
		arg.TicketTypeStatusId,
		arg.ItemId,
		arg.ConditionId,
		arg.RefId,
	)
	var i model.TicketActionConditionRef
	err := row.Scan(
		&i.Uuid,
		&i.TicketTypeStatusId,
		&i.ItemCode,
		&i.ItemId,
		&i.ConditionId,
		&i.RefId,
	)
	return i, err
}

const updateTicketActionConditionRef = `-- name: UpdateTicketActionConditionRef :one
UPDATE Ticket_Action_Condition_Ref SET 
  ticket_type_status_id = $2,
  Item_ID = $3,
  Condition_ID = $4,
  Ref_ID = $5
WHERE uuid = $1
RETURNING UUID, ticket_type_status_id, Item_ID, Condition_Id, Ref_ID
`

func (q *QueriesTransaction) UpdateTicketActionConditionRef(ctx context.Context, arg TicketActionConditionRefRequest) (model.TicketActionConditionRef, error) {
	row := q.db.QueryRowContext(ctx, updateTicketActionConditionRef,
		arg.Uuid,
		arg.TicketTypeStatusId,
		arg.ItemId,
		arg.ConditionId,
		arg.RefId,
	)
	var i model.TicketActionConditionRef
	err := row.Scan(
		&i.Uuid,
		&i.TicketTypeStatusId,
		&i.ItemCode,
		&i.ItemId,
		&i.ConditionId,
		&i.RefId,
	)
	return i, err
}

const deleteTicketActionConditionRef = `-- name: DeleteTicketActionConditionRef :exec
DELETE FROM Ticket_Action_Condition_Ref
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketActionConditionRef(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketActionConditionRef, uuid)
	return err
}

type TicketActionConditionRefInfo struct {
	Uuid               uuid.UUID    `json:"uuid"`
	TicketTypeStatusId int64        `json:"ticketTypeStatusId"`
	ItemCode           string       `json:"itemCode"`
	ItemId             int64        `json:"itemid"`
	Item               string       `json:"item"`
	ItemShortName      string       `json:"itemShortName"`
	ItemDescription    string       `json:"itemDescription"`
	ConditionId        int64        `json:"conditionId"`
	Condition          string       `json:"condition"`
	RefId              int64        `json:"refId"`
	RefTitle           string       `json:"refTitle"`
	RefShortName       string       `json:"refShortName"`
	ModCtr             int64        `json:"mod_ctr"`
	Created            sql.NullTime `json:"created"`
	Updated            sql.NullTime `json:"updated"`
}

func populateTicketActionConditionRef(q *QueriesTransaction, ctx context.Context, sql string) (TicketActionConditionRefInfo, error) {
	var i TicketActionConditionRefInfo
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
		&i.RefId,
		&i.RefTitle,
		&i.RefShortName,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketActionConditionRef2(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketActionConditionRefInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []TicketActionConditionRefInfo{}
	for rows.Next() {
		var i TicketActionConditionRefInfo
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
			&i.RefId,
			&i.RefTitle,
			&i.RefShortName,

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

const TicketActionConditionRefSQL = `-- name: TicketActionConditionRefSQL
SELECT 
  mr.UUID, d.ticket_type_status_id, d.Item_Code, d.Item_ID, i.Title Item, i.Short_Name, i.Remark ItemDescription, 
  Condition_Id, c.Short_Name Condition, d.Ref_Id, ref.Title Ref_Title, ref.Short_Name Ref_Short_Name,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Action_Condition_Ref d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference i on d.Item_Id =  i.Id
INNER JOIN Reference c on d.Condition_Id =  c.Id
INNER JOIN Reference ref on ref.ID = Item_ID
`

func (q *QueriesTransaction) GetTicketActionConditionRef(ctx context.Context, ticketTypeStatusId int64, itemId int64) (TicketActionConditionRefInfo, error) {
	return populateTicketActionConditionRef(q, ctx, fmt.Sprintf("%s WHERE d.Ticket_Type_Status_Id = %v and d.Item_Id = %v",
		TicketActionConditionRefSQL, ticketTypeStatusId, itemId))
}

func (q *QueriesTransaction) GetTicketActionConditionRefbyUuid(ctx context.Context, uuid uuid.UUID) (TicketActionConditionRefInfo, error) {
	return populateTicketActionConditionRef(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", TicketActionConditionRefSQL, uuid))
}

type ListTicketActionConditionRefParams struct {
	TicketTypeStatusId int64 `json:"TicketTypeStatusId"`
	Limit              int32 `json:"limit"`
	Offset             int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketActionConditionRef(ctx context.Context, arg ListTicketActionConditionRefParams) ([]TicketActionConditionRefInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.ticket_type_status_id = %v LIMIT %d OFFSET %d",
			TicketActionConditionRefSQL, arg.TicketTypeStatusId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.ticket_type_status_id = %v ", TicketActionConditionRefSQL, arg.TicketTypeStatusId)
	}
	return populateTicketActionConditionRef2(q, ctx, sql)
}
