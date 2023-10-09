package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createCustomerSpecsString = `-- name: CreateCustomerSpecsString: one
INSERT INTO Customer_Specs_String 
  (Customer_ID, Specs_ID, Value) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Customer_ID, Specs_ID ) DO UPDATE SET
  Value = excluded.Value
RETURNING 
  UUID, Customer_ID, Specs_Code, Specs_ID, Value
`

type CustomerSpecsStringRequest struct {
	Uuid       uuid.UUID `json:"uuid"`
	CustomerId int64     `json:"customerId"`
	SpecsId    int64     `json:"specsId"`
	Value      string    `json:"value"`
}

func (q *QueriesCustomer) CreateCustomerSpecsString(ctx context.Context, arg CustomerSpecsStringRequest) (model.CustomerSpecsString, error) {
	row := q.db.QueryRowContext(ctx, createCustomerSpecsString,
		arg.CustomerId,
		arg.SpecsId,
		arg.Value,
	)
	var i model.CustomerSpecsString
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const updateCustomerSpecsString = `-- name: UpdateCustomerSpecsString :one
UPDATE Customer_Specs_String SET 
Customer_ID = $2,
  Specs_ID = $3,
  Value = $4
WHERE uuid = $1
RETURNING UUID, Customer_ID, Specs_Code, Specs_ID, Value
`

func (q *QueriesCustomer) UpdateCustomerSpecsString(ctx context.Context, arg CustomerSpecsStringRequest) (model.CustomerSpecsString, error) {
	row := q.db.QueryRowContext(ctx, updateCustomerSpecsString,
		arg.Uuid,
		arg.CustomerId,
		arg.SpecsId,
		arg.Value,
	)
	var i model.CustomerSpecsString
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
	)
	return i, err
}

const deleteCustomerSpecsString = `-- name: DeleteCustomerSpecsString :exec
DELETE FROM Customer_Specs_String
WHERE uuid = $1
`

func (q *QueriesCustomer) DeleteCustomerSpecsString(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCustomerSpecsString, uuid)
	return err
}

type CustomerSpecsStringInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	CustomerId      int64        `json:"customerId"`
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

func populateCustomerSpecString(q *QueriesCustomer, ctx context.Context, sql string) (CustomerSpecsStringInfo, error) {
	var i CustomerSpecsStringInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
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

func populateCustomerSpecString2(q *QueriesCustomer, ctx context.Context, sql string) ([]CustomerSpecsStringInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []CustomerSpecsStringInfo{}
	for rows.Next() {
		var i CustomerSpecsStringInfo
		err := rows.Scan(
			&i.Uuid,
			&i.CustomerId,
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

const customerSpecsStringSQL = `-- name: customerSpecsStringSQL
SELECT 
  mr.UUID, d.Customer_ID, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name Item_Short_Name, 
  ref.Remark Item_Description, d.Value,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Specs_String d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesCustomer) GetCustomerSpecsString(ctx context.Context, customerId int64, specsId int64) (CustomerSpecsStringInfo, error) {
	return populateCustomerSpecString(q, ctx, fmt.Sprintf("%s WHERE d.Customer_ID = %v and d.Specs_ID = %v",
		customerSpecsStringSQL, customerId, specsId))
}

func (q *QueriesCustomer) GetCustomerSpecsStringbyUuid(ctx context.Context, uuid uuid.UUID) (CustomerSpecsStringInfo, error) {
	return populateCustomerSpecString(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", customerSpecsStringSQL, uuid))
}

type ListCustomerSpecsStringParams struct {
	CustomerId int64 `json:"CustomerId"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *QueriesCustomer) ListCustomerSpecsString(ctx context.Context, arg ListCustomerSpecsStringParams) ([]CustomerSpecsStringInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Customer_ID = %v LIMIT %d OFFSET %d",
			customerSpecsStringSQL, arg.CustomerId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Customer_ID = %v ", customerSpecsStringSQL, arg.CustomerId)
	}
	return populateCustomerSpecString2(q, ctx, sql)
}
