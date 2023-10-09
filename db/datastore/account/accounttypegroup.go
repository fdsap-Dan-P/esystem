package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createAccountTypeGroup = `-- name: CreateAccountTypeGroup: one
INSERT INTO Account_Type_Group(
   uuid, product_id, group_id, account_type_group, normal_balance, isgl, active, other_info )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (UUID)
DO UPDATE SET
	product_id =  EXCLUDED.product_id,
	group_id =  EXCLUDED.group_id,
	account_type_group =  EXCLUDED.account_type_group,
	normal_balance =  EXCLUDED.normal_balance,
	isgl =  EXCLUDED.isgl,
	active =  EXCLUDED.active
RETURNING id, uuid, product_id, group_id, account_type_group, normal_balance, isgl, active, other_info`

type AccountTypeGroupRequest struct {
	Id               int64     `json:"id"`
	Uuid             uuid.UUID `json:"uuid"`
	ProductId        int64     `json:"productId"`
	GroupId          int64     `json:"groupId"`
	AccountTypeGroup string    `json:"accountTypeGroup"`
	// Iiid             sql.NullInt64  `json:"iiid"`
	NormalBalance bool           `json:"normalBalance"`
	Isgl          bool           `json:"isgl"`
	Active        bool           `json:"active"`
	OtherInfo     sql.NullString `json:"otherInfo"`
}

func (q *QueriesAccount) CreateAccountTypeGroup(ctx context.Context, arg AccountTypeGroupRequest) (model.AccountTypeGroup, error) {
	row := q.db.QueryRowContext(ctx, createAccountTypeGroup,
		arg.Uuid,
		arg.ProductId,
		arg.GroupId,
		arg.AccountTypeGroup,
		// arg.Iiid,
		arg.NormalBalance,
		arg.Isgl,
		arg.Active,
		arg.OtherInfo,
	)
	var i model.AccountTypeGroup
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductId,
		&i.GroupId,
		&i.AccountTypeGroup,
		// &i.Iiid,
		&i.NormalBalance,
		&i.Isgl,
		&i.Active,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccountTypeGroup = `-- name: DeleteAccountTypeGroup :exec
DELETE FROM Account_Type_Group
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteAccountTypeGroup(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountTypeGroup, uuid)
	return err
}

type AccountTypeGroupInfo struct {
	Id               int64     `json:"id"`
	Uuid             uuid.UUID `json:"uuid"`
	ProductId        int64     `json:"productId"`
	GroupId          int64     `json:"groupId"`
	AccountTypeGroup string    `json:"accountTypeGroup"`
	// Iiid             sql.NullInt64  `json:"iiid"`
	NormalBalance bool           `json:"normalBalance"`
	Isgl          bool           `json:"isgl"`
	Active        bool           `json:"active"`
	OtherInfo     sql.NullString `json:"otherInfo"`
	ModCtr        int64          `json:"modCtr"`
	Created       sql.NullTime   `json:"created"`
	Updated       sql.NullTime   `json:"updated"`
}

const accountTypeGroupSQL = `-- name: AccountTypeGroupSQL :one
SELECT
  Id, mr.UUID, product_id, group_id, account_type_group, normal_balance, isgl, active, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Type_Group d 
INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateAccountTypeGroup(q *QueriesAccount, ctx context.Context, sql string) (AccountTypeGroupInfo, error) {
	var i AccountTypeGroupInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductId,
		&i.GroupId,
		&i.AccountTypeGroup,
		// &i.Iiid,
		&i.NormalBalance,
		&i.Isgl,
		&i.Active,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccountTypeGroups(q *QueriesAccount, ctx context.Context, sql string) ([]AccountTypeGroupInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountTypeGroupInfo{}
	for rows.Next() {
		var i AccountTypeGroupInfo
		err := rows.Scan(
			&i.Id,
			&i.Uuid,
			&i.ProductId,
			&i.GroupId,
			&i.AccountTypeGroup,
			// &i.Iiid,
			&i.NormalBalance,
			&i.Isgl,
			&i.Active,
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

func (q *QueriesAccount) GetAccountTypeGroup(ctx context.Context, id int64) (AccountTypeGroupInfo, error) {
	return populateAccountTypeGroup(q, ctx, fmt.Sprintf("%s WHERE d.Id = %v", accountTypeGroupSQL, id))
}

func (q *QueriesAccount) GetAccountTypeGroupbyUuid(ctx context.Context, uuid uuid.UUID) (AccountTypeGroupInfo, error) {
	return populateAccountTypeGroup(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", accountTypeGroupSQL, uuid))
}

type ListAccountTypeGroupParams struct {
	ProductId int64 `json:"productId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountTypeGroup(ctx context.Context, arg ListAccountTypeGroupParams) ([]AccountTypeGroupInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			accountTypeGroupSQL, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf(accountTypeGroupSQL)
	}
	return populateAccountTypeGroups(q, ctx, sql)
}

const updateAccountTypeGroup = `-- name: UpdateAccountTypeGroup :one
UPDATE Account_Type_Group SET 
	product_id = $2,
	group_id = $3,
	account_type_group = $4,
	normal_balance = $5,
	isgl = $6,
	active = $7,
	other_info = $8
WHERE id = $1
RETURNING id, uuid, product_id, group_id, account_type_group, normal_balance, isgl, active, other_info
`

func (q *QueriesAccount) UpdateAccountTypeGroup(ctx context.Context, arg AccountTypeGroupRequest) (model.AccountTypeGroup, error) {
	row := q.db.QueryRowContext(ctx, updateAccountTypeGroup,
		arg.Id,
		arg.ProductId,
		arg.GroupId,
		arg.AccountTypeGroup,
		// arg.Iiid,
		arg.NormalBalance,
		arg.Isgl,
		arg.Active,
		arg.OtherInfo,
	)
	var i model.AccountTypeGroup
	err := row.Scan(
		&i.Id,
		&i.Uuid,
		&i.ProductId,
		&i.GroupId,
		&i.AccountTypeGroup,
		// &i.Iiid,
		&i.NormalBalance,
		&i.Isgl,
		&i.Active,
		&i.OtherInfo,
	)
	return i, err
}
