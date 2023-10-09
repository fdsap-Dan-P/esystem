package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createCustomerBeneficiary = `-- name: CreateCustomerBeneficiary: one
INSERT INTO Customer_Beneficiary (
Customer_Id, Series, IIId, Type_Id, Relation_Type_Id, 
Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6
) 
ON CONFLICT(Customer_Id, IIId) DO UPDATE SET
	Series = excluded.Series,   
	Type_Id = excluded.Type_Id,
	Relation_Type_Id = excluded.Relation_Type_Id,
	Other_Info = excluded.Other_Info
RETURNING UUId, Customer_Id, Series, IIId, Type_Id, Relation_Type_Id, 
Other_Info
`

type CustomerBeneficiaryRequest struct {
	Uuid           uuid.UUID      `json:"uuid"`
	CustomerId     int64          `json:"customerId"`
	Series         int16          `json:"series"`
	Iiid           int64          `json:"iiid"`
	TypeId         int64          `json:"typeId"`
	RelationTypeId int64          `json:"relation_typeId"`
	OtherInfo      sql.NullString `json:"otherInfo"`
}

func (q *QueriesCustomer) CreateCustomerBeneficiary(ctx context.Context, arg CustomerBeneficiaryRequest) (model.CustomerBeneficiary, error) {
	row := q.db.QueryRowContext(ctx, createCustomerBeneficiary,
		arg.CustomerId,
		arg.Series,
		arg.Iiid,
		arg.TypeId,
		arg.RelationTypeId,
		arg.OtherInfo,
	)
	var i model.CustomerBeneficiary
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
		&i.Series,
		&i.Iiid,
		&i.TypeId,
		&i.RelationTypeId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteCustomerBeneficiary = `-- name: DeleteCustomerBeneficiary :exec
DELETE FROM Customer_Beneficiary
WHERE uuid = $1
`

func (q *QueriesCustomer) DeleteCustomerBeneficiary(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCustomerBeneficiary, uuid)
	return err
}

type CustomerBeneficiaryInfo struct {
	Uuid             uuid.UUID      `json:"uuid"`
	CustomerId       int64          `json:"customerId"`
	Series           int16          `json:"series"`
	TitleId          sql.NullInt64  `json:"titleId"`
	LastName         string         `json:"last_name"`
	FirstName        sql.NullString `json:"first_name"`
	MiddleName       sql.NullString `json:"middle_name"`
	MotherMaidenName sql.NullString `json:"mother_maiden_name"`
	Birthday         sql.NullTime   `json:"birthday"`
	Sex              sql.NullBool   `json:"sex"`
	GenderId         sql.NullInt64  `json:"genderId"`
	Iiid             int64          `json:"iiid"`
	TypeId           int64          `json:"typeId"`
	RelationTypeId   int64          `json:"relation_typeId"`
	OtherInfo        sql.NullString `json:"otherInfo"`
	ModCtr           int64          `json:"modCtr"`
	Created          sql.NullTime   `json:"created"`
	Updated          sql.NullTime   `json:"updated"`
}

const getCustomerBeneficiary = `-- name: GetCustomerBeneficiary :one
SELECT 
	mr.UUId, d.Customer_Id, Series,
	ii.Title_Id, ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Mother_Maiden_Name, 
	ii.Birthday, ii.Sex, ii.Gender_id,
	IIId, Type_Id, Relation_Type_Id, d.Other_Info,
	mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Beneficiary d 
INNER JOIN Identity_Info ii on ii.Id = d.IIId
INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesCustomer) GetCustomerBeneficiary(ctx context.Context, uuid uuid.UUID) (CustomerBeneficiaryInfo, error) {
	row := q.db.QueryRowContext(ctx, getCustomerBeneficiary, uuid)
	var i CustomerBeneficiaryInfo
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
		&i.Series,
		&i.TitleId,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.MotherMaidenName,
		&i.Birthday,
		&i.Sex,
		&i.GenderId,
		&i.Iiid,
		&i.TypeId,
		&i.RelationTypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getCustomerBeneficiarybyUuId = `-- name: GetCustomerBeneficiarybyUuId :one
SELECT 
	mr.UUId, d.Customer_Id, Series,
	ii.Title_Id, ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Mother_Maiden_Name, 
	ii.Birthday, ii.Sex, ii.Gender_id,
	IIId, Type_Id, Relation_Type_Id, d.Other_Info,
	mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Beneficiary d 
INNER JOIN Identity_Info ii on ii.Id = d.IIId
INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesCustomer) GetCustomerBeneficiarybyUuId(ctx context.Context, uuid uuid.UUID) (CustomerBeneficiaryInfo, error) {
	row := q.db.QueryRowContext(ctx, getCustomerBeneficiarybyUuId, uuid)
	var i CustomerBeneficiaryInfo
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
		&i.Series,
		&i.TitleId,
		&i.LastName,
		&i.FirstName,
		&i.MiddleName,
		&i.MotherMaidenName,
		&i.Birthday,
		&i.Sex,
		&i.GenderId,
		&i.Iiid,
		&i.TypeId,
		&i.RelationTypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listCustomerBeneficiary = `-- name: ListCustomerBeneficiary:many
SELECT 
	mr.UUId, d.Customer_Id, Series,
	ii.Title_Id, ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Mother_Maiden_Name, 
	ii.Birthday, ii.Sex, ii.Gender_id,
	IIId, Type_Id, Relation_Type_Id, d.Other_Info,
	mr.Mod_Ctr, mr.Created, mr.Updated
FROM Customer_Beneficiary d 
INNER JOIN Identity_Info ii on ii.Id = d.IIId
INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Customer_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListCustomerBeneficiaryParams struct {
	CustomerId int64 `json:"customerId"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *QueriesCustomer) ListCustomerBeneficiary(ctx context.Context, arg ListCustomerBeneficiaryParams) ([]CustomerBeneficiaryInfo, error) {
	rows, err := q.db.QueryContext(ctx, listCustomerBeneficiary, arg.CustomerId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CustomerBeneficiaryInfo{}
	for rows.Next() {
		var i CustomerBeneficiaryInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.CustomerId,
			&i.Series,
			&i.TitleId,
			&i.LastName,
			&i.FirstName,
			&i.MiddleName,
			&i.MotherMaidenName,
			&i.Birthday,
			&i.Sex,
			&i.GenderId,
			&i.Iiid,
			&i.TypeId,
			&i.RelationTypeId,
			&i.OtherInfo,

			&i.ModCtr,
			&i.Created,
			&i.Updated,
		); err != nil {
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

const updateCustomerBeneficiary = `-- name: UpdateCustomerBeneficiary :one
UPDATE Customer_Beneficiary SET 
Customer_Id = $2,
Series = $3,
IIId = $4,
Type_Id = $5,
Relation_Type_Id = $6,
Other_Info = $7
WHERE uuid = $1
RETURNING UUId, Customer_Id, Series, IIId, Type_Id, Relation_Type_Id, 
Other_Info
`

func (q *QueriesCustomer) UpdateCustomerBeneficiary(ctx context.Context, arg CustomerBeneficiaryRequest) (model.CustomerBeneficiary, error) {
	row := q.db.QueryRowContext(ctx, updateCustomerBeneficiary,

		arg.Uuid,
		arg.CustomerId,
		arg.Series,
		arg.Iiid,
		arg.TypeId,
		arg.RelationTypeId,
		arg.OtherInfo,
	)
	var i model.CustomerBeneficiary
	err := row.Scan(
		&i.Uuid,
		&i.CustomerId,
		&i.Series,
		&i.Iiid,
		&i.TypeId,
		&i.RelationTypeId,
		&i.OtherInfo,
	)
	return i, err
}
