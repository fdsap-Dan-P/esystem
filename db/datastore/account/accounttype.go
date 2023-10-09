package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"simplebank/model"
)

const createAccountType = `-- name: CreateAccountType: one
INSERT INTO Account_Type(
	uuid, central_office_id, code, account_type, product_id,
	group_id, normal_balance, isgl, active, filter_type, other_info )
 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
 RETURNING 
   id, uuid, central_office_id, code, account_type, product_id, group_id, 
   normal_balance, isgl, active, filter_type, other_info

`

type AccountTypeRequest struct {
	Id              int64     `json:"id"`
	Uuid            uuid.UUID `json:"uuid"`
	CentralOfficeId int64     `json:"centralOfficeId"`
	Code            int64     `json:"code"`
	AccountType     string    `json:"accountType"`
	ProductId       int64     `json:"productId"`
	GroupId         int64     `json:"groupId"`
	// Iiid            sql.NullInt64  `json:"iiid"`
	NormalBalance bool           `json:"normalBalance"`
	Isgl          bool           `json:"isgl"`
	Active        bool           `json:"active"`
	FilterType    int64          `json:"filterType"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}

func (q *QueriesAccount) CreateAccountType(ctx context.Context, arg AccountTypeRequest) (model.AccountType, error) {
	row := q.db.QueryRowContext(ctx, createAccountType,
		arg.Uuid,
		arg.CentralOfficeId,
		arg.Code,
		arg.AccountType,
		arg.ProductId,
		arg.GroupId,
		// arg.Iiid,
		arg.NormalBalance,
		arg.Isgl,
		arg.Active,
		arg.FilterType,
		arg.OtherInfo,
	)
	var i model.AccountType
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.Code,
		&i.AccountType,
		&i.ProductId,
		&i.GroupId,
		// &i.Iiid,
		&i.NormalBalance,
		&i.Isgl,
		&i.Active,
		&i.FilterType,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccountType = `-- name: DeleteAccountType :exec
	DELETE FROM Account_Type
	WHERE id = $1
	`

func (q *QueriesAccount) DeleteAccountType(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccountType, id)
	return err
}

type AccountTypeInfo struct {
	Id              int64     `json:"id"`
	Uuid            uuid.UUID `json:"uuid"`
	CentralOfficeId int64     `json:"centralOfficeId"`
	Code            int64     `json:"code"`
	AccountType     string    `json:"accountType"`
	ProductId       int64     `json:"productId"`
	// ProductName     string    `json:"productName"`
	GroupId int64 `json:"groupId"`
	// Iiid            sql.NullInt64  `json:"iiid"`
	NormalBalance bool           `json:"normalBalance"`
	Isgl          bool           `json:"isgl"`
	Active        bool           `json:"active"`
	FilterType    int64          `json:"filterType"`
	OtherInfo     sql.NullString `json:"otherInfo"`
	ModCtr        int64          `json:"modCtr"`
	Created       sql.NullTime   `json:"created"`
	Updated       sql.NullTime   `json:"updated"`
}

const accountTypeSQL = `-- name: AccountTypeSQL :one
	SELECT
	ID, mr.UUID, central_office_id, code, account_type, product_id, group_id, normal_balance, isgl, active, filter_type, other_info
	,mr.Mod_Ctr, mr.Created, mr.Updated
	FROM Account_Type d INNER JOIN Main_Record mr on mr.UUID = d.UUID
	`

func populateAccountType(q *QueriesAccount, ctx context.Context, sql string) (AccountTypeInfo, error) {
	var i AccountTypeInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.Code,
		&i.AccountType,
		&i.ProductId,
		&i.GroupId,
		// &i.Iiid,
		&i.NormalBalance,
		&i.Isgl,
		&i.Active,
		&i.FilterType,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccountTypes(q *QueriesAccount, ctx context.Context, sql string) ([]AccountTypeInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountTypeInfo{}
	for rows.Next() {
		var i AccountTypeInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.CentralOfficeId,
			&i.Code,
			&i.AccountType,
			&i.ProductId,
			&i.GroupId,
			// &i.Iiid,
			&i.NormalBalance,
			&i.Isgl,
			&i.Active,
			&i.FilterType,
			&i.OtherInfo,
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

func (q *QueriesAccount) GetAccountType(ctx context.Context, id int64) (AccountTypeInfo, error) {
	return populateAccountType(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", accountTypeSQL, id))
}

func (q *QueriesAccount) GetAccountTypebyUuid(ctx context.Context, uuid uuid.UUID) (AccountTypeInfo, error) {
	return populateAccountType(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", accountTypeSQL, uuid))
}

func (q *QueriesAccount) GetAccountTypebyName(ctx context.Context, actType string) (AccountTypeInfo, error) {
	return populateAccountType(q, ctx, fmt.Sprintf("%s WHERE lower(d.Account_Type) = lower('%v')", accountTypeSQL, actType))
}

type ListAccountTypeParams struct {
	ProductId int64 `json:"productId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountType(ctx context.Context, arg ListAccountTypeParams) ([]AccountTypeInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			accountTypeSQL, arg.Limit, arg.Offset)
	} else {
		sql = accountTypeSQL
	}
	return populateAccountTypes(q, ctx, sql)
}

const updateAccountType = `-- name: UpdateAccountType :one
UPDATE Account_Type SET 
	uuid = $2,
	central_office_id = $3,
	code = $4,
	account_type = $5,
	product_id = $6,
	group_id = $7,
	normal_balance = $8,
	isgl = $9,
	active = $10,
	filter_type = $11,
	other_info = $12
WHERE id = $1
RETURNING id, uuid, central_office_id, code, account_type, product_id, group_id, normal_balance, isgl, active, filter_type, other_info
	`

func (q *QueriesAccount) UpdateAccountType(ctx context.Context, arg AccountTypeRequest) (model.AccountType, error) {
	row := q.db.QueryRowContext(ctx, updateAccountType,
		arg.Id,
		arg.Uuid,
		arg.CentralOfficeId,
		arg.Code,
		arg.AccountType,
		arg.ProductId,
		arg.GroupId,
		// arg.Iiid,
		arg.NormalBalance,
		arg.Isgl,
		arg.Active,
		arg.FilterType,
		arg.OtherInfo,
	)
	var i model.AccountType
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.CentralOfficeId,
		&i.Code,
		&i.AccountType,
		&i.ProductId,
		&i.GroupId,
		// &i.Iiid,
		&i.NormalBalance,
		&i.Isgl,
		&i.Active,
		&i.FilterType,
		&i.OtherInfo,
	)
	return i, err
}
