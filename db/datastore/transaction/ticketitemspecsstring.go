package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createTicketItemSpecsString = `-- name: CreateTicketItemSpecsString: one
INSERT INTO Ticket_Item_Specs_String 
  (Ticket_Item_ID, Specs_ID, Value) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Ticket_Item_ID, Specs_ID ) DO UPDATE SET
  Value = excluded.Value
RETURNING 
  UUID, Ticket_Item_ID, Specs_Code, Specs_ID, Value
`

type TicketItemSpecsStringRequest struct {
	Uuid         uuid.UUID `json:"uuid"`
	TicketItemId int64     `json:"TicketItemId"`
	SpecsId      int64     `json:"specsId"`
	Value        string    `json:"value"`
}

func (q *QueriesTransaction) CreateTicketItemSpecsString(ctx context.Context, arg TicketItemSpecsStringRequest) (model.TicketItemSpecsString, error) {
	row := q.db.QueryRowContext(ctx, createTicketItemSpecsString,
		arg.TicketItemId,
		arg.SpecsId,
		arg.Value,
	)
	var i model.TicketItemSpecsString
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const updateTicketItemSpecsString = `-- name: UpdateTicketItemSpecsString :one
UPDATE Ticket_Item_Specs_String SET 
Ticket_Item_ID = $2,
  Specs_ID = $3,
  Value = $4
WHERE uuid = $1
RETURNING UUID, Ticket_Item_ID, Specs_Code, Specs_ID, Value
`

func (q *QueriesTransaction) UpdateTicketItemSpecsString(ctx context.Context, arg TicketItemSpecsStringRequest) (model.TicketItemSpecsString, error) {
	row := q.db.QueryRowContext(ctx, updateTicketItemSpecsString,
		arg.Uuid,
		arg.TicketItemId,
		arg.SpecsId,
		arg.Value,
	)
	var i model.TicketItemSpecsString
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const deleteTicketItemSpecsString = `-- name: DeleteTicketItemSpecsString :exec
DELETE FROM Ticket_Item_Specs_String
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketItemSpecsString(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketItemSpecsString, uuid)
	return err
}

type TicketItemSpecsStringInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	TicketItemId    int64        `json:"TicketItemId"`
	SpecsCode       string       `json:"specsCode"`
	SpecsId         int64        `json:"specsId"`
	Item            string       `json:"item"`
	ItemShortName   string       `json:"itemShortName"`
	ItemDescription string       `json:"itemDescription"`
	Value           string       `json:"value"`
	ModCtr          int64        `json:"mod_ctr"`
	Created         sql.NullTime `json:"created"`
	Updated         sql.NullTime `json:"updated"`
}

func populateTicketItemSpecString(q *QueriesTransaction, ctx context.Context, sql string) (TicketItemSpecsStringInfo, error) {
	var i TicketItemSpecsStringInfo
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

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketItemSpecString2(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketItemSpecsStringInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []TicketItemSpecsStringInfo{}
	for rows.Next() {
		var i TicketItemSpecsStringInfo
		err := rows.Scan(
			&i.Uuid,
			&i.TicketItemId,
			&i.SpecsCode,
			&i.SpecsId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
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

const TicketItemSpecsStringSQL = `-- name: TicketItemSpecsStringSQL
SELECT 
  mr.UUID, d.Ticket_Item_ID, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name Item_Short_Name, 
  ref.Remark Item_Description, d.Value,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Item_Specs_String d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesTransaction) GetTicketItemSpecsString(ctx context.Context, TicketItemId int64, specsId int64) (TicketItemSpecsStringInfo, error) {
	return populateTicketItemSpecString(q, ctx, fmt.Sprintf("%s WHERE d.Ticket_Item_ID = %v and d.Specs_ID = %v",
		TicketItemSpecsStringSQL, TicketItemId, specsId))
}

func (q *QueriesTransaction) GetTicketItemSpecsStringbyUuid(ctx context.Context, uuid uuid.UUID) (TicketItemSpecsStringInfo, error) {
	return populateTicketItemSpecString(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", TicketItemSpecsStringSQL, uuid))
}

type ListTicketItemSpecsStringParams struct {
	TicketItemId int64 `json:"TicketItemId"`
	Limit        int32 `json:"limit"`
	Offset       int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketItemSpecsString(ctx context.Context, arg ListTicketItemSpecsStringParams) ([]TicketItemSpecsStringInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Ticket_Item_ID = %v LIMIT %d OFFSET %d",
			TicketItemSpecsStringSQL, arg.TicketItemId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Ticket_Item_ID = %v ", TicketItemSpecsStringSQL, arg.TicketItemId)
	}
	return populateTicketItemSpecString2(q, ctx, sql)
}
