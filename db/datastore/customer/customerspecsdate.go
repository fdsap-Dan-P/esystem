package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"simplebank/model"
)

const createCustomerSpecsDate = `-- name: CreateCustomerSpecsDate: one
INSERT INTO Customer_Specs_Date 
  (Customer_Id, Specs_ID, Value, Value2) 
VALUES 
  ($1, $2, $3, $4) 
ON CONFLICT( Customer_Id, Specs_ID ) DO UPDATE SET
  Value = excluded.Value,
  Value2 = excluded.Value2
RETURNING 
  UUID, Customer_Id, Specs_Code, Specs_ID, Value, Value2
`

type CustomerSpecsDateRequest struct {
	Uuid       uuid.UUID `json:"uuid"`
	CustomerId int64     `json:"customerId"`
	SpecsId    int64     `json:"specsId"`
	Value      time.Time `json:"value"`
	Value2     time.Time `json:"value2"`
}

func (q *QueriesCustomer) CreateCustomerSpecsDate(ctx context.Context, arg CustomerSpecsDateRequest) (model.CustomerSpecsDate, error) {
	row := q.db.QueryRowContext(ctx, createCustomerSpecsDate,
		arg.CustomerId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.CustomerSpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const updateCustomerSpecsDate = `-- name: UpdateCustomerSpecsDate :one
UPDATE Customer_Specs_Date SET 
Customer_Id = $2,
	Specs_ID = $3,
	Value = $4,
	Value2 = $5
WHERE uuid = $1
RETURNING UUID, Customer_Id, Specs_Code, Specs_ID, Value, Value2
`

func (q *QueriesCustomer) UpdateCustomerSpecsDate(ctx context.Context, arg CustomerSpecsDateRequest) (model.CustomerSpecsDate, error) {
	row := q.db.QueryRowContext(ctx, updateCustomerSpecsDate,
		arg.Uuid,
		arg.CustomerId,
		arg.SpecsId,
		arg.Value,
		arg.Value2,
	)
	var i model.CustomerSpecsDate
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.Value,
		&i.Value2,
	)
	return i, err
}

const deleteCustomerSpecsDate = `-- name: DeleteCustomerSpecsDate :exec
DELETE FROM Customer_Specs_Date
WHERE uuid = $1
`

func (q *QueriesCustomer) DeleteCustomerSpecsDate(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCustomerSpecsDate, uuid)
	return err
}

type CustomerSpecsDateInfo struct {
	Uuid            uuid.UUID    `json:"uuid"`
	CustomerId      int64        `json:"customerId"`
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

func populateIdentitySpecDate(q *QueriesCustomer, ctx context.Context, sql string) (CustomerSpecsDateInfo, error) {
	var i CustomerSpecsDateInfo
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
		&i.Value2,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateIdentitySpecDate2(q *QueriesCustomer, ctx context.Context, sql string) ([]CustomerSpecsDateInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []CustomerSpecsDateInfo{}
	for rows.Next() {
		var i CustomerSpecsDateInfo
		err := rows.Scan(
			&i.Uuid,
			&i.CustomerId,
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

const CustomerSpecsDateSQL = `-- name: CustomerSpecsDateSQL
SELECT 
  mr.UUID, d.Customer_ID, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, 
  d.Value, d.Value2,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Specs_Date d INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesCustomer) GetCustomerSpecsDate(ctx context.Context, customerId int64, specsId int64) (CustomerSpecsDateInfo, error) {
	return populateIdentitySpecDate(q, ctx, fmt.Sprintf("%s WHERE d.Customer_ID = %v and d.Specs_ID = %v",
		CustomerSpecsDateSQL, customerId, specsId))
}

func (q *QueriesCustomer) GetCustomerSpecsDatebyUuid(ctx context.Context, uuid uuid.UUID) (CustomerSpecsDateInfo, error) {
	return populateIdentitySpecDate(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", CustomerSpecsDateSQL, uuid))
}

type ListCustomerSpecsDateParams struct {
	CustomerId int64 `json:"CustomerId"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *QueriesCustomer) ListCustomerSpecsDate(ctx context.Context, arg ListCustomerSpecsDateParams) ([]CustomerSpecsDateInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Customer_ID = %v LIMIT %d OFFSET %d",
			CustomerSpecsDateSQL, arg.CustomerId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Customer_ID = %v ", CustomerSpecsDateSQL, arg.CustomerId)
	}
	return populateIdentitySpecDate2(q, ctx, sql)
}
