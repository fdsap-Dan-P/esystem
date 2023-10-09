package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
)

const createTicketItemSpecsDate = `-- name: CreateTicketItemSpecsDate: one
INSERT INTO Ticket_Item_Specs_Date 
  (Ticket_Item_Id, Specs_ID, Value, Value2) 
VALUES 
  ($1, $2, $3, $4) 
ON CONFLICT( Ticket_Item_Id, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Ticket_Item_Id, Specs_Code, Specs_ID, Value, Value2
`

type TicketItemSpecsDateRequest struct {
	Uuid         uuid.UUID `json:"uuid"`
	TicketItemId int64     `json:"TicketItemId"`
	SpecsId      int64     `json:"specsId"`
	Value        time.Time `json:"value"`
	Value2       time.Time `json:"value2"`
}

func (q *QueriesTransaction) CreateTicketItemSpecsDate(ctx context.Context, arg TicketItemSpecsDateRequest) (model.TicketItemSpecsDate, error) {
	row := q.db.QueryRowContext(ctx, createTicketItemSpecsDate,
		arg.TicketItemId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.TicketItemSpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const updateTicketItemSpecsDate = `-- name: UpdateTicketItemSpecsDate :one
UPDATE Ticket_Item_Specs_Date SET 
Ticket_Item_Id = $2,
	Specs_ID = $3,
	Value = $4,
	Value2 = $5
WHERE uuid = $1
RETURNING UUID, Ticket_Item_Id, Specs_Code, Specs_ID, Value, Value2
`

func (q *QueriesTransaction) UpdateTicketItemSpecsDate(ctx context.Context, arg TicketItemSpecsDateRequest) (model.TicketItemSpecsDate, error) {
	row := q.db.QueryRowContext(ctx, updateTicketItemSpecsDate,
		arg.Uuid,
		arg.TicketItemId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.TicketItemSpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const deleteTicketItemSpecsDate = `-- name: DeleteTicketItemSpecsDate :exec
DELETE FROM Ticket_Item_Specs_Date
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketItemSpecsDate(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketItemSpecsDate, uuid)
	return err
}

type TicketItemSpecsDateInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	TicketItemId    int64        `json:"TicketItemId"`
	SpecsCode       string       `json:"specsCode"`
	SpecsId         int64        `json:"specsId"`
	Item            string       `json:"item"`
	ItemShortName   string       `json:"itemShortName"`
	ItemDescription string       `json:"itemDescription"`
	Value           time.Time    `json:"value"`
	Value2          time.Time    `json:"value2"`
	ModCtr          int64        `json:"modCtr"`
	Created         sql.NullTime `json:"created"`
	Updated         sql.NullTime `json:"updated"`
}

func populateTicketItemSpecDate(q *QueriesTransaction, ctx context.Context, sql string) (TicketItemSpecsDateInfo, error) {
	var i TicketItemSpecsDateInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Item,
		&i.ItemShortName,
		&i.ItemDescription,
		&i.Value,
		&i.Value2,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketItemSpecDate2(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketItemSpecsDateInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []TicketItemSpecsDateInfo{}
	for rows.Next() {
		var i TicketItemSpecsDateInfo
		err := rows.Scan(
			&i.Uuid,
			&i.TicketItemId,
			&i.SpecsCode,
			&i.SpecsId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
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

const TicketItemSpecsDateSQL = `-- name: TicketItemSpecsDateSQL
SELECT 
  mr.UUID, d.Ticket_Item_Id, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Item_Specs_Date d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesTransaction) GetTicketItemSpecsDate(ctx context.Context, TicketItemId int64, specsId int64) (TicketItemSpecsDateInfo, error) {
	return populateTicketItemSpecDate(q, ctx, fmt.Sprintf("%s WHERE d.Ticket_Item_Id = %v and d.Specs_ID = %v",
		TicketItemSpecsDateSQL, TicketItemId, specsId))
}

func (q *QueriesTransaction) GetTicketItemSpecsDatebyUuid(ctx context.Context, uuid uuid.UUID) (TicketItemSpecsDateInfo, error) {
	return populateTicketItemSpecDate(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", TicketItemSpecsDateSQL, uuid))
}

type ListTicketItemSpecsDateParams struct {
	TicketItemId int64 `json:"TicketItemId"`
	Limit        int32 `json:"limit"`
	Offset       int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketItemSpecsDate(ctx context.Context, arg ListTicketItemSpecsDateParams) ([]TicketItemSpecsDateInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Ticket_Item_Id = %v LIMIT %d OFFSET %d",
			TicketItemSpecsDateSQL, arg.TicketItemId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Ticket_Item_Id = %v ", TicketItemSpecsDateSQL, arg.TicketItemId)
	}
	return populateTicketItemSpecDate2(q, ctx, sql)
}
