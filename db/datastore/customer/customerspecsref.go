package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createCustomerSpecsRef = `-- name: CreateCustomerSpecsRef: one
INSERT INTO Customer_Specs_Ref 
  (Customer_Id, Specs_ID, Ref_Id) 
VALUES 
  ($1, $2, $3) 
ON CONFLICT( Customer_Id, Specs_ID ) DO UPDATE SET
Ref_Id = excluded.Ref_Id
RETURNING 
  UUID, Customer_Id, Specs_Code, Specs_ID, Ref_Id
`

type CustomerSpecsRefRequest struct {
	Uuid       uuid.UUID `json:"uuid"`
	CustomerId int64     `json:"customerId"`
	SpecsId    int64     `json:"specsId"`
	RefId      int64     `json:"refId"`
}

func (q *QueriesCustomer) CreateCustomerSpecsRef(ctx context.Context, arg CustomerSpecsRefRequest) (model.CustomerSpecsRef, error) {
	row := q.db.QueryRowContext(ctx, createCustomerSpecsRef,
		arg.CustomerId,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.CustomerSpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const updateCustomerSpecsRef = `-- name: UpdateCustomerSpecsRef :one
UPDATE Customer_Specs_Ref SET 
  Customer_Id = $2,
  Specs_ID = $3,
  Ref_ID = $4
WHERE uuid = $1
RETURNING UUID, Customer_Id, Specs_Code, Specs_ID, Ref_ID
`

func (q *QueriesCustomer) UpdateCustomerSpecsRef(ctx context.Context, arg CustomerSpecsRefRequest) (model.CustomerSpecsRef, error) {
	row := q.db.QueryRowContext(ctx, updateCustomerSpecsRef,
		arg.Uuid,
		arg.CustomerId,
		arg.SpecsId,
		arg.RefId,
	)
	var i model.CustomerSpecsRef
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
		&i.SpecsCode,
		&i.SpecsId,
		&i.RefId,
	)
	return i, err
}

const deleteCustomerSpecsRef = `-- name: DeleteCustomerSpecsRef :exec
DELETE FROM Customer_Specs_Ref
WHERE uuid = $1
`

func (q *QueriesCustomer) DeleteCustomerSpecsRef(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCustomerSpecsRef, uuid)
	return err
}

type CustomerSpecsRefInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	CustomerId      int64          `json:"customerId"`
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

func populateCustomerSpecRef(q *QueriesCustomer, ctx context.Context, sql string) (CustomerSpecsRefInfo, error) {
	var i CustomerSpecsRefInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
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

func populateCustomerSpecRef2(q *QueriesCustomer, ctx context.Context, sql string) ([]CustomerSpecsRefInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := []CustomerSpecsRefInfo{}
	for rows.Next() {
		var i CustomerSpecsRefInfo
		err := rows.Scan(
			&i.Uuid,
			&i.CustomerId,
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

const customerSpecsRefSQL = `-- name: customerSpecsRefSQL
SELECT 
  mr.UUID, d.Customer_Id, d.Specs_Code, d.Specs_ID, ref.Title, ref.Short_Name, ref.Remark ItemDescription, d.Ref_Id,
  mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Specs_Ref d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
INNER JOIN Reference ref on ref.ID = Specs_ID
`

func (q *QueriesCustomer) GetCustomerSpecsRef(ctx context.Context, customerId int64, specsId int64) (CustomerSpecsRefInfo, error) {
	return populateCustomerSpecRef(q, ctx, fmt.Sprintf("%s WHERE d.Customer_Id = %v and d.Specs_ID = %v",
		customerSpecsRefSQL, customerId, specsId))
}

func (q *QueriesCustomer) GetCustomerSpecsRefbyUuid(ctx context.Context, uuid uuid.UUID) (CustomerSpecsRefInfo, error) {
	return populateCustomerSpecRef(q, ctx, fmt.Sprintf("%s WHERE mr.UUID = '%v'", customerSpecsRefSQL, uuid))
}

type ListCustomerSpecsRefParams struct {
	CustomerId int64 `json:"CustomerId"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *QueriesCustomer) ListCustomerSpecsRef(ctx context.Context, arg ListCustomerSpecsRefParams) ([]CustomerSpecsRefInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s WHERE d.Customer_Id = %v LIMIT %d OFFSET %d",
			customerSpecsRefSQL, arg.CustomerId, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf("%s WHERE d.Customer_Id = %v ", customerSpecsRefSQL, arg.CustomerId)
	}
	return populateCustomerSpecRef2(q, ctx, sql)
}
