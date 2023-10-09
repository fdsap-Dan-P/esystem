package db

import (
	"context"
	"database/sql"
	"fmt"
	"simplebank/model"

	"github.com/google/uuid"
)

const createAccountTypeFilter = `-- name: CreateAccountTypeFilter: one
INSERT INTO Account_Type_Filter(
	UUID, Account_Type_Id, Central_Office_Id, allow, other_info )
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT(UUID)
DO UPDATE SET
    Account_Type_Id = EXCLUDED.Account_Type_Id,
	Central_Office_Id = EXCLUDED.Central_Office_Id,
	Allow = EXCLUDED.Allow,
	other_info = EXCLUDED.other_info
RETURNING UUID, Account_Type_Id, Central_Office_Id, allow, other_info`

type AccountTypeFilterRequest struct {
	Uuid            uuid.UUID      `json:"uuid"`
	AccountTypeId   int64          `json:"accountTypeId"`
	CentralOfficeId int64          `json:"centralOfficeId"`
	Allow           bool           `json:"allow"`
	OtherInfo       sql.NullString `json:"otherInfo"`
}

func (q *QueriesAccount) CreateAccountTypeFilter(ctx context.Context, arg AccountTypeFilterRequest) (model.AccountTypeFilter, error) {
	row := q.db.QueryRowContext(ctx, createAccountTypeFilter,
		arg.Uuid,
		arg.AccountTypeId,
		arg.CentralOfficeId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.AccountTypeFilter
	err := row.Scan(
		&i.Uuid,
		&i.AccountTypeId,
		&i.CentralOfficeId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}

const deleteAccountTypeFilter = `-- name: DeleteAccountTypeFilter :exec
DELETE FROM Account_Type_Filter
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteAccountTypeFilter(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccountTypeFilter, uuid)
	return err
}

type AccountTypeFilterInfo struct {
	Uuid            uuid.UUID      `json:"uuid"`
	AccountTypeId   int64          `json:"accountTypeId"`
	CentralOfficeId int64          `json:"centralOfficeId"`
	Allow           bool           `json:"allow"`
	OtherInfo       sql.NullString `json:"otherInfo"`
	ModCtr          int64          `json:"modCtr"`
	Created         sql.NullTime   `json:"created"`
	Updated         sql.NullTime   `json:"updated"`
}

const accountTypeFilterSQL = `-- name: AccountTypeFilterSQL :one
SELECT
  mr.UUID, Account_Type_Id, Central_Office_Id, allow, other_info
  ,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Account_Type_Filter d INNER JOIN Main_Record mr on mr.UUID = d.UUID
`

func populateAccountTypeFilter(q *QueriesAccount, ctx context.Context, sql string) (AccountTypeFilterInfo, error) {
	var i AccountTypeFilterInfo
	row := q.db.QueryRowContext(ctx, sql)
	err := row.Scan(
		&i.Uuid,
		&i.AccountTypeId,
		&i.CentralOfficeId,
		&i.Allow,
		&i.OtherInfo,
		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

func populateAccountTypeFilters(q *QueriesAccount, ctx context.Context, sql string) ([]AccountTypeFilterInfo, error) {
	rows, err := q.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountTypeFilterInfo{}
	for rows.Next() {
		var i AccountTypeFilterInfo
		err := rows.Scan(
			&i.Uuid,
			&i.AccountTypeId,
			&i.CentralOfficeId,
			&i.Allow,
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

func (q *QueriesAccount) GetAccountTypeFilter(ctx context.Context, officeId int64, acctTypeID int64) (AccountTypeFilterInfo, error) {
	return populateAccountTypeFilter(q, ctx, fmt.Sprintf("%s WHERE d.Central_Office_Id = %v and Account_Type_Id = %v", accountTypeFilterSQL, officeId, acctTypeID))
}

func (q *QueriesAccount) GetAccountTypeFilterbyUuid(ctx context.Context, uuid uuid.UUID) (AccountTypeFilterInfo, error) {
	return populateAccountTypeFilter(q, ctx, fmt.Sprintf("%s WHERE d.UUID = '%v'", accountTypeFilterSQL, uuid))
}

type ListAccountTypeFilterParams struct {
	CentralOfficeId int64 `json:"centralOfficeId"`
	Limit           int32 `json:"limit"`
	Offset          int32 `json:"offset"`
}

func (q *QueriesAccount) ListAccountTypeFilter(ctx context.Context, arg ListAccountTypeFilterParams) ([]AccountTypeFilterInfo, error) {
	var sql string
	if arg.Limit != 0 {
		sql = fmt.Sprintf("%s  LIMIT %d OFFSET %d",
			accountTypeFilterSQL, arg.Limit, arg.Offset)
	} else {
		sql = fmt.Sprintf(accountTypeFilterSQL)
	}
	return populateAccountTypeFilters(q, ctx, sql)
}

const updateAccountTypeFilter = `-- name: UpdateAccountTypeFilter :one
UPDATE Account_Type_Filter SET 
    Account_Type_Id = $2,
	Central_Office_Id = $3,
	Allow = $4,
	other_info = $5
WHERE uuid = $1
RETURNING uuid, account_type_id, central_office_id, allow, other_info
`

func (q *QueriesAccount) UpdateAccountTypeFilter(ctx context.Context, arg AccountTypeFilterRequest) (model.AccountTypeFilter, error) {
	row := q.db.QueryRowContext(ctx, updateAccountTypeFilter,
		arg.Uuid,
		arg.AccountTypeId,
		arg.CentralOfficeId,
		arg.Allow,
		arg.OtherInfo,
	)
	var i model.AccountTypeFilter
	err := row.Scan(
		&i.Uuid,
		&i.AccountTypeId,
		&i.CentralOfficeId,
		&i.Allow,
		&i.OtherInfo,
	)
	return i, err
}
