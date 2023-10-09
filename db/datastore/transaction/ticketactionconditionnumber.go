package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createTicketActionConditionNumber = `-- name: CreateTicketActionConditionNumber: one
INSERT INTO Ticket_Action_Condition_Number 
  (uuid, ticket_type_status_id, Item_ID, Condition_Id, Value, Value2, Measure_Id) 
VALUES 
  ($1, $2, $3, $4, $5, $6, $7) 
ON CONFLICT(uuid) DO UPDATE SET
  ticket_type_status_id = excluded.ticket_type_status_id, 
  Item_ID = excluded.Item_ID,
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, ticket_type_status_id, Item_ID, Condition_Id, Value, Value2, Measure_Id
`

type TicketActionConditionNumberRequest struct {
	Uuid               uuid.UUID       `json:"uuid"`
	TicketTypeStatusId int64           `json:"TicketTypeStatusId"`
	ItemId             int64           `json:"itemid"`
	ConditionId        int64           `json:"conditionId"`
	Value              decimal.Decimal `json:"value"`
	Value2             decimal.Decimal `json:"value2"`
	MeasureId          sql.NullInt64   `json:"measureId"`
}

func (q *QueriesTransaction) CreateTicketActionConditionNumber(ctx context.Context, arg TicketActionConditionNumberRequest) (model.TicketActionConditionNumber, error) {
	row := q.db.QueryRowContext(ctx, createTicketActionConditionNumber,
		arg.Uuid,
		arg.TicketTypeStatusId,
		arg.ItemId,
		arg.ConditionId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.TicketActionConditionNumber
	err := row.Scan(
		&i.Uuid,
		&i.TicketTypeStatusId,
		&i.ItemCode,
		&i.ItemId,
		&i.ConditionId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const updateTicketActionConditionNumber = `-- name: UpdateTicketActionConditionNumber :one
UPDATE Ticket_Action_Condition_Number SET 
  ticket_type_status_id = $2,
  Item_ID = $3,
  Condition_ID = $4,
  Value = $5,
  Value2 = $6,
  Measure_Id = $7
WHERE uuid = $1
RETURNING UUID, ticket_type_status_id, Item_ID, Condition_ID, Value, Value2, Measure_Id
`

func (q *QueriesTransaction) UpdateTicketActionConditionNumber(ctx context.Context, arg TicketActionConditionNumberRequest) (model.TicketActionConditionNumber, error) {
	row := q.db.QueryRowContext(ctx, updateTicketActionConditionNumber,
		arg.Uuid,
		arg.TicketTypeStatusId,
		arg.ItemId,
		arg.ConditionId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.TicketActionConditionNumber
	err := row.Scan(
		&i.Uuid,
		&i.TicketTypeStatusId,
		&i.ItemCode,
		&i.ItemId,
		&i.ConditionId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const deleteTicketActionConditionNumber = `-- name: DeleteTicketActionConditionNumber :exec
DELETE FROM Ticket_Action_Condition_Number
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketActionConditionNumber(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketActionConditionNumber, uuid)
	return err
}

type TicketActionConditionNumberInfo struct {
	Uuid               uuid.UUID       `json:"uuid"`
	TicketTypeStatusId int64           `json:"TicketTypeStatusId"`
	ItemId             int64           `json:"itemid"`
	ItemCode           string          `json:"ItemCode"`
	Item               string          `json:"item"`
	ItemShortName      string          `json:"itemShortName"`
	ItemDescription    string          `json:"itemDescription"`
	ConditionId        int64           `json:"conditionId"`
	Condition          string          `json:"condition"`
	Value              decimal.Decimal `json:"value"`
	Value2             decimal.Decimal `json:"value2"`
	MeasureId          sql.NullInt64   `json:"measureId"`
	Measure            sql.NullString  `json:"measure"`
	MeasureUnit        sql.NullString  `json:"measureUnit"`
	ModCtr             int64           `json:"mod_ctr"`
	Created            sql.NullTime    `json:"created"`
	Updated            sql.NullTime    `json:"updated"`
}

func populateTicketActionConfitionNumber(q *QueriesTransaction, ctx context.Context, sql string) (TicketActionConditionNumberInfo, error) {
	var i TicketActionConditionNumberInfo
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
		&i.Value2,
		&i.MeasureId,
		&i.Measure,
		&i.MeasureUnit,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketActionConfitionNumber2(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketActionConditionNumberInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []TicketActionConditionNumberInfo{}
	for rows.Next() {
		var i TicketActionConditionNumberInfo
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
			&i.Value2,
			&i.MeasureId,
			&i.Measure,
			&i.MeasureUnit,

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

const TicketActionConditionNumberSQL = `-- name: TicketActionConditionNumberSQL
SELECT 
  mr.UUID, d.ticket_type_status_id, item_id, i.Title Item, i.Short_Name, i.Remark ItemDescription, 
  condition_id, c.Short_Name Condition, 
  d.Value, d.Value2, d.Measure_Id, mea.Title Measure, mea.Short_Name Measure_Unit, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Action_Condition_Number d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference i on d.Item_Id =  i.Id
INNER JOIN Reference c on d.Condition_Id =  c.Id
LEFT JOIN Reference mea on mea.ID = Measure_Id
`

func (q *QueriesTransaction) GetTicketActionConditionNumber(ctx context.Context, ticketTypeStatusId int64, itemId int64) (TicketActionConditionNumberInfo, error) {
	log.Printf("ticketTypeStatusId : %v itemId: %v", ticketTypeStatusId, itemId)
	return populateTicketActionConfitionNumber(q, ctx, fmt.Sprintf("%s WHERE d.Ticket_Type_Status_Id = %v and d.Item_Id = %v",
		TicketActionConditionNumberSQL, ticketTypeStatusId, itemId))
}

func (q *QueriesTransaction) GetTicketActionConditionNumberbyUuid(ctx context.Context, uuid uuid.UUID) (TicketActionConditionNumberInfo, error) {
	return populateTicketActionConfitionNumber(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", TicketActionConditionNumberSQL, uuid))
}

type ListTicketActionConditionNumberParams struct {
	TicketTypeStatusId int64 `json:"TicketTypeStatusId"`
	Limit              int32 `json:"limit"`
	Offset             int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketActionConditionNumber(ctx context.Context, arg ListTicketActionConditionNumberParams) ([]TicketActionConditionNumberInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.ticket_type_status_id = %v LIMIT %d OFFSET %d",
			TicketActionConditionNumberSQL, arg.TicketTypeStatusId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.ticket_type_status_id = %v ", TicketActionConditionNumberSQL, arg.TicketTypeStatusId)
	}
	return populateTicketActionConfitionNumber2(q, ctx, sql)
}
