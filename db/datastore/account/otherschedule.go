package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"simplebank/model"

	"github.com/shopspring/decimal"
)

const createOtherSchedule = `-- name: CreateOtherSchedule: one
INSERT INTO Other_Schedule (
Account_Id, Charge_Id, Series, Due_Date, Due_Amt, 
Realizable, End_Bal, Other_Info
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8
) 
ON CONFLICT(Account_Id, Charge_Id, Series) DO UPDATE SET
Account_Id = excluded.Account_Id,
Charge_Id = excluded.Charge_Id,
Series = excluded.Series,
Due_Date  = excluded.Due_Date,
Due_Amt  = excluded.Due_Amt,
Realizable = excluded. Realizable,
End_Bal = excluded.End_Bal

RETURNING UUId, Account_Id, Charge_Id, Series, Due_Date, Due_Amt, 
Realizable, End_Bal, Other_Info
`

type OtherScheduleRequest struct {
	Uuid       uuid.UUID       `json:"uuid"`
	AccountId  int64           `json:"accountId"`
	ChargeId   int64           `json:"chargeId"`
	Series     int16           `json:"series"`
	DueDate    time.Time       `json:"dueDate"`
	DueAmt     decimal.Decimal `json:"dueAmt"`
	Realizable decimal.Decimal `json:"realizable"`
	EndBal     decimal.Decimal `json:"endBal"`
	OtherInfo  sql.NullString  `json:"otherInfo"`
}

func (q *QueriesAccount) CreateOtherSchedule(ctx context.Context, arg OtherScheduleRequest) (model.OtherSchedule, error) {
	row := q.db.QueryRowContext(ctx, createOtherSchedule,
		arg.AccountId,
		arg.ChargeId,
		arg.Series,
		arg.DueDate,
		arg.DueAmt,
		arg.Realizable,
		arg.EndBal,
		arg.OtherInfo,
	)
	var i model.OtherSchedule
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.ChargeId,
		&i.Series,
		&i.DueDate,
		&i.DueAmt,
		&i.Realizable,
		&i.EndBal,
		&i.OtherInfo,
	)
	return i, err
}

const deleteOtherSchedule = `-- name: DeleteOtherSchedule :exec
DELETE FROM Other_Schedule
WHERE uuid = $1
`

func (q *QueriesAccount) DeleteOtherSchedule(ctx context.Context, uuid uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteOtherSchedule, uuid)
	return err
}

type OtherScheduleInfo struct {
	Uuid       uuid.UUID       `json:"uuid"`
	AccountId  int64           `json:"accountId"`
	ChargeId   int64           `json:"chargeId"`
	Series     int16           `json:"series"`
	DueDate    time.Time       `json:"dueDate"`
	DueAmt     decimal.Decimal `json:"dueAmt"`
	Realizable decimal.Decimal `json:"realizable"`
	EndBal     decimal.Decimal `json:"endBal"`
	OtherInfo  sql.NullString  `json:"otherInfo"`

	ModCtr  int64        `json:"modCtr"`
	Created sql.NullTime `json:"created"`
	Updated sql.NullTime `json:"updated"`
}

const getOtherSchedule = `-- name: GetOtherSchedule :one
SELECT 
mr.UUId, Account_Id, Charge_Id, Series, 
Due_Date, Due_Amt, Realizable, End_Bal, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Other_Schedule d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.uuid = $1 LIMIT 1
`

func (q *QueriesAccount) GetOtherSchedule(ctx context.Context, uuid uuid.UUID) (OtherScheduleInfo, error) {
	row := q.db.QueryRowContext(ctx, getOtherSchedule, uuid)
	var i OtherScheduleInfo
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.ChargeId,
		&i.Series,
		&i.DueDate,
		&i.DueAmt,
		&i.Realizable,
		&i.EndBal,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getOtherSchedulebyUuid = `-- name: GetOtherSchedulebyUuid :one
SELECT 
mr.UUId, Account_Id, Charge_Id, Series, 
Due_Date, Due_Amt, Realizable, End_Bal, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Other_Schedule d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE mr.UUId = $1 LIMIT 1
`

func (q *QueriesAccount) GetOtherSchedulebyUuid(ctx context.Context, uuid uuid.UUID) (OtherScheduleInfo, error) {
	row := q.db.QueryRowContext(ctx, getOtherSchedulebyUuid, uuid)
	var i OtherScheduleInfo
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.ChargeId,
		&i.Series,
		&i.DueDate,
		&i.DueAmt,
		&i.Realizable,
		&i.EndBal,
		&i.OtherInfo,

		&i.ModCtr,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listOtherSchedule = `-- name: ListOtherSchedule:many
SELECT 
mr.UUId, Account_Id, Charge_Id, Series, 
Due_Date, Due_Amt, Realizable, End_Bal, Other_Info
,mr.Mod_Ctr, mr.Created, mr.Updated
FROM Other_Schedule d INNER JOIN Main_Record mr on mr.UUId = d.UUId
WHERE Account_Id = $1
ORDER BY uuid
LIMIT $2
OFFSET $3
`

type ListOtherScheduleParams struct {
	AccountId int64 `json:"AccountId"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *QueriesAccount) ListOtherSchedule(ctx context.Context, arg ListOtherScheduleParams) ([]OtherScheduleInfo, error) {
	rows, err := q.db.QueryContext(ctx, listOtherSchedule, arg.AccountId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OtherScheduleInfo{}
	for rows.Next() {
		var i OtherScheduleInfo
		if err := rows.Scan(
			&i.Uuid,
			&i.AccountId,
			&i.ChargeId,
			&i.Series,
			&i.DueDate,
			&i.DueAmt,
			&i.Realizable,
			&i.EndBal,
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

const updateOtherSchedule = `-- name: UpdateOtherSchedule :one
UPDATE Other_Schedule SET 
Account_Id = $2,
Charge_Id = $3,
Series = $4,
Due_Date = $5,
Due_Amt = $6,
Realizable = $7,
End_Bal = $8,
Other_Info = $9
WHERE uuid = $1
RETURNING UUId, Account_Id, Charge_Id, Series, Due_Date, Due_Amt, 
Realizable, End_Bal, Other_Info
`

func (q *QueriesAccount) UpdateOtherSchedule(ctx context.Context, arg OtherScheduleRequest) (model.OtherSchedule, error) {
	row := q.db.QueryRowContext(ctx, updateOtherSchedule,

		arg.Uuid,
		arg.AccountId,
		arg.ChargeId,
		arg.Series,
		arg.DueDate,
		arg.DueAmt,
		arg.Realizable,
		arg.EndBal,
		arg.OtherInfo,
	)
	var i model.OtherSchedule
	err := row.Scan(
		&i.Uuid,
		&i.AccountId,
		&i.ChargeId,
		&i.Series,
		&i.DueDate,
		&i.DueAmt,
		&i.Realizable,
		&i.EndBal,
		&i.OtherInfo,
	)
	return i, err
}
