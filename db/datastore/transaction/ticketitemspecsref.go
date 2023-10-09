package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createTicketItemSpecsRef = `-- name: CreateTicketItemSpecsRef: one
INSERT INTO Ticket_Item_Specs_Ref 
  (Ticket_Item_ID, Specs_ID, Ref_Id) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Ticket_Item_ID, Specs_ID ) DO UPDATE SET
Ref_Id = excluded.Ref_Id
RETURNING 
  UUID, Ticket_Item_ID, Specs_Code, Specs_ID, Ref_Id
`

type TicketItemSpecsRefRequest struct {
	Uuid         uuid.UUID `json:"uuid"`
	TicketItemId int64     `json:"TicketItemId"`
	SpecsId      int64     `json:"specsId"`
	RefId        int64     `json:"refId"`
}

func (q *QueriesTransaction) CreateTicketItemSpecsRef(ctx context.Context, arg TicketItemSpecsRefRequest) (model.TicketItemSpecsRef, error) {
	row := q.db.QueryRowContext(ctx, createTicketItemSpecsRef,
		arg.TicketItemId,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.TicketItemSpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const updateTicketItemSpecsRef = `-- name: UpdateTicketItemSpecsRef :one
UPDATE Ticket_Item_Specs_Ref SET 
  Ticket_Item_ID = $2,
  Specs_ID = $3,
  Ref_ID = $4
WHERE uuid = $1
RETURNING UUID, Ticket_Item_ID, Specs_Code, Specs_ID, Ref_ID
`

func (q *QueriesTransaction) UpdateTicketItemSpecsRef(ctx context.Context, arg TicketItemSpecsRefRequest) (model.TicketItemSpecsRef, error) {
	row := q.db.QueryRowContext(ctx, updateTicketItemSpecsRef,
		arg.Uuid,
		arg.TicketItemId,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.TicketItemSpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const deleteTicketItemSpecsRef = `-- name: DeleteTicketItemSpecsRef :exec
DELETE FROM Ticket_Item_Specs_Ref
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketItemSpecsRef(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketItemSpecsRef, uuid)
	return err
}

type TicketItemSpecsRefInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	TicketItemId    int64          `json:"TicketItemId"`
	SpecsCode       string         `json:"specsCode"`
	SpecsId         int64          `json:"specsId"`
	Item            string         `json:"item"`
	ItemShortName   string         `json:"itemShortName"`
	ItemDescription string         `json:"itemDescription"`
	RefId           int64          `json:"refId"`
	MeasureId       sql.NullInt64  `json:"measureId"`
	Measure         sql.NullString `json:"measure"`
	MeasureUnit     sql.NullString `json:"measureUnit"`
	ModCtr          int64          `json:"mod_ctr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

func populateTicketItemSpecRef(q *QueriesTransaction, ctx context.Context, sql string) (TicketItemSpecsRefInfo, error) {
	var i TicketItemSpecsRefInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Item,
		&i.ItemShortName,
		&i.ItemDescription,
		&i.RefId,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketItemSpecRef2(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketItemSpecsRefInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []TicketItemSpecsRefInfo{}
	for rows.Next() {
		var i TicketItemSpecsRefInfo
		err := rows.Scan(
			&i.Uuid,
			&i.TicketItemId,
			&i.SpecsCode,
			&i.SpecsId,
			&i.Item,
			&i.ItemShortName,
			&i.ItemDescription,
			&i.RefId,

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

const TicketItemSpecsRefSQL = `-- name: TicketItemSpecsRefSQL
SELECT 
  mr.UUID, d.Ticket_Item_ID, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, d.Ref_Id,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Item_Specs_Ref d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesTransaction) GetTicketItemSpecsRef(ctx context.Context, TicketItemId int64, specsId int64) (TicketItemSpecsRefInfo, error) {
	return populateTicketItemSpecRef(q, ctx, fmt.Sprintf("%s WHERE d.Ticket_Item_ID = %v and d.Specs_ID = %v",
		TicketItemSpecsRefSQL, TicketItemId, specsId))
}

func (q *QueriesTransaction) GetTicketItemSpecsRefbyUuid(ctx context.Context, uuid uuid.UUID) (TicketItemSpecsRefInfo, error) {
	return populateTicketItemSpecRef(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", TicketItemSpecsRefSQL, uuid))
}

type ListTicketItemSpecsRefParams struct {
	TicketItemId int64 `json:"TicketItemId"`
	Limit        int32 `json:"limit"`
	Offset       int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketItemSpecsRef(ctx context.Context, arg ListTicketItemSpecsRefParams) ([]TicketItemSpecsRefInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Ticket_Item_ID = %v LIMIT %d OFFSET %d",
			TicketItemSpecsRefSQL, arg.TicketItemId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Ticket_Item_ID = %v ", TicketItemSpecsRefSQL, arg.TicketItemId)
	}
	return populateTicketItemSpecRef2(q, ctx, sql)
}
