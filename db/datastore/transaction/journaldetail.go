package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createJournalDetail = `-- name: CreateJournalDetail: one
INSERT INTO Journal_Detail (
Trn_Head_Id, Series, Office_Id, COA_Id, Account_Type_Id, 
Currency, Partition_Id, Trn_Amt, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING UUId, Trn_Head_Id, Series, Office_Id, COA_Id, Account_Type_Id, 
Currency, Partition_Id, Trn_Amt, Other_Info
`

type JournalDetailRequest struct {
	Uuid          uuid.UUID       `json:"uuid"`
	TrnHeadId     int64           `json:"trnHeadId"`
	Series        int16           `json:"series"`
	OfficeId      int64           `json:"officeId"`
	CoaId         sql.NullInt64   `json:"coaId"`
	AccountTypeId sql.NullInt64   `json:"accountTypeId"`
	Currency      sql.NullString  `json:"currency"`
	PartitionId   sql.NullInt64   `json:"partitionId"`
	TrnAmt        decimal.Decimal `json:"trn_amt"`
	OtherInfo     sql.NullString  `json:"otherInfo"`
}

func (q *QueriesTransaction) CreateJournalDetail(ctx context.Context, arg JournalDetailRequest) (model.JournalDetail, error) {
	row := q.db.QueryRowContext(ctx, createJournalDetail,
		arg.TrnHeadId,
		arg.Series,
		arg.OfficeId,
		arg.CoaId,
		arg.AccountTypeId,
		arg.Currency,
		arg.PartitionId,
		arg.TrnAmt,
		arg.OtherInfo,
	)
	var i model.JournalDetail
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.OfficeId,
		&i.CoaId,
		&i.AccountTypeId,
		&i.Currency,
		&i.PartitionId,
		&i.TrnAmt,
		&i.OtherInfo,
	)
	return i, err
}

const deleteJournalDetail = `-- name: DeleteJournalDetail :exec
DELETE FROM Journal_Detail
WHERE uuid = $1
`

func (q *QueriesTransaction) DeleteJournalDetail(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteJournalDetail, uuid)
	return err
}

type JournalDetailInfo struct {
	Uuid          uuid.UUID       `json:"uuid"`
	TrnHeadId     int64           `json:"trnHeadId"`
	Series        int16           `json:"series"`
	OfficeId      int64           `json:"officeId"`
	CoaId         sql.NullInt64   `json:"coaId"`
	AccountTypeId sql.NullInt64   `json:"accountTypeId"`
	Currency      sql.NullString  `json:"currency"`
	PartitionId   sql.NullInt64   `json:"partitionId"`
	TrnAmt        decimal.Decimal `json:"trn_amt"`
	OtherInfo     sql.NullString  `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getJournalDetail = `-- name: GetJournalDetail :one
SELECT 
mr.UUId, Trn_Head_Id, Series, Office_Id, COA_Id, 
Account_Type_Id, Currency, Partition_Id, Trn_Amt, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Journal_Detail d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesTransaction) GetJournalDetail(ctx context.Context, uuid uuid.UUID) (JournalDetailInfo, error) {
	row := q.db.QueryRowContext(ctx, getJournalDetail, uuid)
	var i JournalDetailInfo
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.OfficeId,
		&i.CoaId,
		&i.AccountTypeId,
		&i.Currency,
		&i.PartitionId,
		&i.TrnAmt,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getJournalDetailbyUuid = `-- name: GetJournalDetailbyUuid :one
SELECT 
mr.UUId, Trn_Head_Id, Series, Office_Id, COA_Id, 
Account_Type_Id, Currency, Partition_Id, Trn_Amt, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Journal_Detail d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesTransaction) GetJournalDetailbyUuid(ctx context.Context, uuid uuid.UUID) (JournalDetailInfo, error) {
	row := q.db.QueryRowContext(ctx, getJournalDetailbyUuid, uuid)
	var i JournalDetailInfo
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.OfficeId,
		&i.CoaId,
		&i.AccountTypeId,
		&i.Currency,
		&i.PartitionId,
		&i.TrnAmt,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listJournalDetail = `-- name: ListJournalDetail:many
SELECT 
mr.UUId, Trn_Head_Id, Series, Office_Id, COA_Id, 
Account_Type_Id, Currency, Partition_Id, Trn_Amt, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Journal_Detail d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE trn_head_id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListJournalDetailParams struct {
	TrnHeadId int64 `json:"trnHeadId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesTransaction) ListJournalDetail(ctx context.Context, arg ListJournalDetailParams) ([]JournalDetailInfo, error) {
	rows, err := q.db.QueryContext(ctx, listJournalDetail, arg.TrnHeadId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []JournalDetailInfo{}
	for rows.Next() {
		var i JournalDetailInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.TrnHeadId,
			&i.Series,
			&i.OfficeId,
			&i.CoaId,
			&i.AccountTypeId,
			&i.Currency,
			&i.PartitionId,
			&i.TrnAmt,
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

const updateJournalDetail = `-- name: UpdateJournalDetail :one
UPDATE Journal_Detail SET 
	Trn_Head_Id = $2,
	Series = $3,
	Office_Id = $4,
	COA_Id = $5,
	Account_Type_Id = $6,
	Currency = $7,
	Partition_Id = $8,
	Trn_Amt = $9,
	Other_Info = $10
WHERE uuid = $1
RETURNING UUId, Trn_Head_Id, Series, Office_Id, COA_Id, Account_Type_Id, 
Currency, Partition_Id, Trn_Amt, Other_Info
`

func (q *QueriesTransaction) UpdateJournalDetail(ctx context.Context, arg JournalDetailRequest) (model.JournalDetail, error) {
	row := q.db.QueryRowContext(ctx, updateJournalDetail,

		arg.Uuid,
		arg.TrnHeadId,
		arg.Series,
		arg.OfficeId,
		arg.CoaId,
		arg.AccountTypeId,
		arg.Currency,
		arg.PartitionId,
		arg.TrnAmt,
		arg.OtherInfo,
	)
	var i model.JournalDetail
	err := row.Scan(
		&i.Uuid,
		&i.TrnHeadId,
		&i.Series,
		&i.OfficeId,
		&i.CoaId,
		&i.AccountTypeId,
		&i.Currency,
		&i.PartitionId,
		&i.TrnAmt,
		&i.OtherInfo,
	)
	return i, err
}
