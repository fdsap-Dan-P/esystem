package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccountBeneficiary = `-- name: CreateAccountBeneficiary: one
INSERT INTO Account_Beneficiary (
Account_Id, Series, Beneficiary_Type_Id, IIID, Relationship_Type_Id, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6) 
ON CONFLICT(Account_Id, IIID)
DO UPDATE SET
  Series = EXCLUDED.Series,
  Beneficiary_Type_Id = EXCLUDED.Beneficiary_Type_Id,
  Relationship_Type_Id = EXCLUDED.Relationship_Type_Id,
  Other_Info = EXCLUDED.Other_Info
RETURNING 
  UUID, Account_Id, Series, Beneficiary_Type_Id, IIId, Relationship_Type_Id, Other_Info
`

type AccountBeneficiaryRequest struct {
	Uuid               uuid.UUID      `json:"uuid"`
	AccountId          int64          `json:"accountId"`
	Iiid               int64          `json:"iiid"`
	Series             int16          `json:"series"`
	BeneficiaryTypeId  int64          `json:"beneficiaryTypeId"`
	RelationshipTypeId int64          `json:"relationshipTypeId"`
	OtherInfo          sql.NullString `json:"otherInfo"`
}

func (q *QueriesAccount) CreateAccountBeneficiary(ctx context.Context, arg AccountBeneficiaryRequest) (model.AccountBeneficiary, error) {
	row := q.db.QueryRowContext(ctx, createAccountBeneficiary,
		arg.AccountId,
		arg.Series,
		arg.BeneficiaryTypeId,
		arg.Iiid,
		arg.RelationshipTypeId,
		arg.OtherInfo,
	)
	var i model.AccountBeneficiary
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.Series,
		&i.BeneficiaryTypeId,
		&i.Iiid,
		&i.RelationshipTypeId,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccountBeneficiary = `-- name: DeleteAccountBeneficiary :exec
DELETE FROM Account_Beneficiary
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteAccountBeneficiary(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountBeneficiary, uuid)
	return err
}

type AccountBeneficiaryInfo struct {
	Uuid               uuid.UUID      `json:"uuid"`
	AccountId          int64          `json:"accountId"`
	Iiid               int64          `json:"iiid"`
	Series             int16          `json:"series"`
	BeneficiaryTypeId  int64          `json:"beneficiaryTypeId"`
	RelationshipTypeId int64          `json:"relationshipTypeId"`
	OtherInfo          sql.NullString `json:"otherInfo"`
	ModCtr             int64          `json:"modCtr"`
	Created            sql.NullTime   `json:"created"`
	Updated            sql.NullTime   `json:"updated"`
}

const getAccountBeneficiary = `-- name: GetAccountBeneficiary :one
SELECT 
mr.UUId, Account_Id, 
Series, Beneficiary_Type_Id, IIId, Relationship_Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Beneficiary d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesAccount) GetAccountBeneficiary(ctx context.Context, uuid uuid.UUID) (AccountBeneficiaryInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccountBeneficiary, uuid)
	var i AccountBeneficiaryInfo
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.Series,
		&i.BeneficiaryTypeId,
		&i.Iiid,
		&i.RelationshipTypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAccountBeneficiarybyUuid = `-- name: GetAccountBeneficiarybyUuid :one
SELECT 
mr.UUId, Account_Id, 
Series, Beneficiary_Type_Id, IIId, Relationship_Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Beneficiary d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccount) GetAccountBeneficiarybyUuid(ctx context.Context, uuid uuid.UUID) (AccountBeneficiaryInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccountBeneficiarybyUuid, uuid)
	var i AccountBeneficiaryInfo
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.Series,
		&i.BeneficiaryTypeId,
		&i.Iiid,
		&i.RelationshipTypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getAccountBeneficiarybyName = `-- name: GetAccountBeneficiarybyName :one
SELECT 
mr.UUId, Account_Id, 
Series, Beneficiary_Type_Id, IIId, Relationship_Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Beneficiary d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Title = $1 LIMIT 1
`

func (q *QueriesAccount) GetAccountBeneficiarybyName(ctx context.Context, name string) (AccountBeneficiaryInfo, error) {
	row := q.db.QueryRowContext(ctx, getAccountBeneficiarybyName, name)
	var i AccountBeneficiaryInfo
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.Series,
		&i.BeneficiaryTypeId,
		&i.Iiid,
		&i.RelationshipTypeId,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listAccountBeneficiary = `-- name: ListAccountBeneficiary:many
SELECT 
mr.UUId, Account_Id, 
Series, Beneficiary_Type_Id, IIId, Relationship_Type_Id, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Beneficiary d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Account_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListAccountBeneficiaryParams struct {
	AccountId int64 `json:"accountId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountBeneficiary(ctx context.Context, arg ListAccountBeneficiaryParams) ([]AccountBeneficiaryInfo, error) {
	rows, err := q.db.QueryContext(ctx, listAccountBeneficiary, arg.AccountId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountBeneficiaryInfo{}
	for rows.Next() {
		var i AccountBeneficiaryInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.AccountId,
			&i.Series,
			&i.BeneficiaryTypeId,
			&i.Iiid,
			&i.RelationshipTypeId,
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

const updateAccountBeneficiary = `-- name: UpdateAccountBeneficiary :one
UPDATE Account_Beneficiary SET 
Account_Id = $2,
Series = $3,
Beneficiary_Type_Id = $4,
IIID = $5,
Relationship_Type_Id = $6,
Other_Info = $7
WHERE uuid = $1
RETURNING UUId, Account_Id, Series, Beneficiary_Type_Id, IIId, Relationship_Type_Id, 
Other_Info
`

func (q *QueriesAccount) UpdateAccountBeneficiary(ctx context.Context, arg AccountBeneficiaryRequest) (model.AccountBeneficiary, error) {
	row := q.db.QueryRowContext(ctx, updateAccountBeneficiary,

		arg.Uuid,
		arg.AccountId,
		arg.Series,
		arg.BeneficiaryTypeId,
		arg.Iiid,
		arg.RelationshipTypeId,
		arg.OtherInfo,
	)
	var i model.AccountBeneficiary
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.Series,
		&i.BeneficiaryTypeId,
		&i.Iiid,
		&i.RelationshipTypeId,
		&i.OtherInfo,
	)
	return i, err
}
