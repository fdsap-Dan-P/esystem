package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"simplebank/model"
)

const createTicketItemSpecsNumber = `-- name: CreateTicketItemSpecsNumber: one
INSERT INTO Ticket_Item_Specs_Number 
  (Ticket_Item_ID, Specs_ID, Value, Value2, Measure_Id) 
VALUES 
  ($1, $2, $3, $4, $5) 
ON CONFLICT( Ticket_Item_ID, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Ticket_Item_ID, Specs_Code, Specs_ID, Value, Value2, Measure_Id
`

type TicketItemSpecsNumberRequest struct {
	Uuid         uuid.UUID       `json:"uuid"`
	TicketItemId int64           `json:"TicketItemId"`
	SpecsId      int64           `json:"specsId"`
	Value        decimal.Decimal `json:"value"`
	Value2       decimal.Decimal `json:"value2"`
	MeasureId    sql.NullInt64   `json:"measureId"`
}

func (q *QueriesTransaction) CreateTicketItemSpecsNumber(ctx context.Context, arg TicketItemSpecsNumberRequest) (model.TicketItemSpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, createTicketItemSpecsNumber,
		arg.TicketItemId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.TicketItemSpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		arg.MeasureId,
	)
	return i, err
}

const updateTicketItemSpecsNumber = `-- name: UpdateTicketItemSpecsNumber :one
UPDATE Ticket_Item_Specs_Number SET 
  Ticket_Item_ID = $2,
  Specs_ID = $3,
  Value = $4,
  Value2 = $5,
  Measure_Id = $6
WHERE uuid = $1
RETURNING UUID, Ticket_Item_ID, Specs_Code, Specs_ID, Value, Value2, Measure_Id
`

func (q *QueriesTransaction) UpdateTicketItemSpecsNumber(ctx context.Context, arg TicketItemSpecsNumberRequest) (model.TicketItemSpecsNumber, error) {
	row := q.db.QueryRowContext(ctx, updateTicketItemSpecsNumber,
		arg.Uuid,
		arg.TicketItemId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
		arg.MeasureId,
	)
	var i model.TicketItemSpecsNumber
	err := row.Scan(
		&i.Uuid,
		&i.TicketItemId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
		&i.MeasureId,
	)
	return i, err
}

const deleteTicketItemSpecsNumber = `-- name: DeleteTicketItemSpecsNumber :exec
DELETE FROM Ticket_Item_Specs_Number
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteTicketItemSpecsNumber(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicketItemSpecsNumber, uuid)
	return err
}

type TicketItemSpecsNumberInfo struct {
	Uuid            uuid.UUID       `json:"uuid"`
	TicketItemId    int64           `json:"TicketItemId"`
	SpecsId         int64           `json:"specsId"`
	SpecsCode       string          `json:"specsCode"`
	Item            string          `json:"item"`
	ItemShortName   string          `json:"itemShortName"`
	ItemDescription string          `json:"itemDescription"`
	Value           decimal.Decimal `json:"value"`
	Value2          decimal.Decimal `json:"value2"`
	MeasureId       sql.NullInt64   `json:"measureId"`
	Measure         sql.NullString  `json:"measure"`
	MeasureUnit     sql.NullString  `json:"measureUnit"`
	ModCtr          int64           `json:"mod_ctr"`
	Created         sql.NullTime    `json:"created"`
	Updated         sql.NullTime    `json:"updated"`
}

func populateTicketItemSpecNumber(q *QueriesTransaction, ctx context.Context, sql string) (TicketItemSpecsNumberInfo, error) {
	var i TicketItemSpecsNumberInfo
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
		&i.MeasureId,
		&i.Measure,
		&i.MeasureUnit,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateTicketItemSpecNumber2(q *QueriesTransaction, ctx context.Context, sql string) ([]TicketItemSpecsNumberInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []TicketItemSpecsNumberInfo{}
	for rows.Next() {
		var i TicketItemSpecsNumberInfo
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

const TicketItemSpecsNumberSQL = `-- name: TicketItemSpecsNumberSQL
SELECT 
  mr.UUID, d.Ticket_Item_ID, d.Specs_Code, d.Specs_ID, ref.Title Item, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2, d.Measure_Id, mea.Title Measure, mea.Short_Name Measure_Unit, 
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Ticket_Item_Specs_Number d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
LEFT JOIN Reference mea on mea.ID = Measure_Id`

func (q *QueriesTransaction) GetTicketItemSpecsNumber(ctx context.Context, TicketItemId int64, specsId int64) (TicketItemSpecsNumberInfo, error) {
	return populateTicketItemSpecNumber(q, ctx, fmt.Sprintf("%s WHERE d.Ticket_Item_ID = %v and d.Specs_ID = %v",
		TicketItemSpecsNumberSQL, TicketItemId, specsId))
}

func (q *QueriesTransaction) GetTicketItemSpecsNumberbyUuid(ctx context.Context, uuid uuid.UUID) (TicketItemSpecsNumberInfo, error) {
	return populateTicketItemSpecNumber(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", TicketItemSpecsNumberSQL, uuid))
}

type ListTicketItemSpecsNumberParams struct {
	TicketItemId int64 `json:"TicketItemId"`
	Limit        int32 `json:"limit"`
	Offset       int32 `json:"offset"`
}

func (q *QueriesTransaction) ListTicketItemSpecsNumber(ctx context.Context, arg ListTicketItemSpecsNumberParams) ([]TicketItemSpecsNumberInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Ticket_Item_ID = %v LIMIT %d OFFSET %d",
			TicketItemSpecsNumberSQL, arg.TicketItemId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Ticket_Item_ID = %v ", TicketItemSpecsNumberSQL, arg.TicketItemId)
	}
	return populateTicketItemSpecNumber2(q, ctx, sql)
}
